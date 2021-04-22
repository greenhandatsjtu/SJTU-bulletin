# -*- coding: utf-8 -*-
import re

import scrapy

from ..items import NoticeItem


class NoticeSpider(scrapy.Spider):
    name = 'notice'

    # 声明请求链接和对应的解析函数
    def start_requests(self):
        yield scrapy.Request(url='http://www.jwc.sjtu.edu.cn/xwtg/tztg.htm', callback=self.parseNewJwc)
        yield scrapy.Request(url='http://www.jwc.sjtu.edu.cn/index/mxxsdtz.htm', callback=self.parseNewJwc)
        yield scrapy.Request(url='http://xsb.seiee.sjtu.edu.cn/xsb/list/705-1-20.htm', callback=self.parseXsb)
        yield scrapy.Request(url='http://xsb.seiee.sjtu.edu.cn/xsb/list/3016-1-20.htm', callback=self.parseXsb)
        yield scrapy.Request(url='http://xsb.seiee.sjtu.edu.cn/xsb/list/2496-1-20.htm', callback=self.parsePartTime)
        yield scrapy.Request(url='http://xsb.seiee.sjtu.edu.cn/xsb/list/2495-1-20.htm', callback=self.parseFullTime)
        yield scrapy.Request(url='https://www.sjtu.edu.cn/tg/index.html', callback=self.parseSjtuNotice)
        yield scrapy.Request(url='http://ourhome.sjtu.edu.cn/news', callback=self.parseOurHome)

    # 爬取教务处通知
    def parseJwc(self, response):
        # 利用css selector得到所需信息
        news = response.css('.main_r_xuxian tr')
        for new in news:
            item = NoticeItem()
            item['title'] = new.css('a::text').get()
            item['date'] = new.css("td:nth-child(2)::text").get().lstrip("\r\n\t\t\t\t [").rstrip(" ]\r\n\t\t\t ")
            item['href'] = response.urljoin(new.css('a::attr(href)').get())
            item['_type'] = "jwc"
            yield item

    # 爬取新教务处通知
    def parseNewJwc(self, response):
        news = response.css('.Newslist li.clearfix')
        for new in news:
            item = NoticeItem()
            item['title'] = new.css('.wz>a>h2::text').get()
            item['href'] = response.urljoin(new.css('.wz>a::attr(href)').get())
            item['_type'] = "jwc"
            # yield item
            yield scrapy.Request(item['href'], callback=self.getJwcNoticeDate, meta={"item": item})

    # get jwc notice publish date
    def getJwcNoticeDate(self, response):
        item = response.request.meta['item']
        item['date'] = response.css('.content-title>i::text').re(r'\d{4}-\d{2}-\d{2}')[0]
        yield item

    # 爬取学生办通知
    def parseXsb(self, response):
        news = response.css('.list_box_5_2>li')
        for new in news:
            item = NoticeItem()
            item['title'] = new.css('a::attr(title)').extract_first().lstrip("<b>").rstrip("</b>").strip()
            item['date'] = new.css('span::text').extract_first().lstrip("[").rstrip("]")
            item['href'] = response.urljoin(new.css('a::attr("href")').extract_first())
            item['_type'] = "xsb"
            yield item

    # 爬取实习招聘信息
    def parsePartTime(self, response):
        news = response.css('.list_box_5_2>li')
        for new in news:
            item = NoticeItem()
            item['title'] = new.css('a::attr(title)').extract_first().lstrip("<b>").rstrip("</b>").strip()
            item['date'] = new.css('span::text').extract_first().lstrip("[").rstrip("]")
            item['href'] = response.urljoin(new.css('a::attr("href")').extract_first())
            item['_type'] = "partTime"
            yield item

    # 爬取全职招聘信息
    def parseFullTime(self, response):
        news = response.css('.list_box_5_2>li')
        for new in news:
            item = NoticeItem()
            item['title'] = new.css('a::attr(title)').extract_first().lstrip("<b>").rstrip("</b>").strip()
            item['date'] = new.css('span::text').extract_first().lstrip("[").rstrip("]")
            item['href'] = response.urljoin(new.css('a::attr("href")').extract_first())
            item['_type'] = "fullTime"
            yield item

    # 爬取交大官网通知通告
    def parseSjtuNotice(self, response):
        news = response.css('.pageMain li')
        for new in news:
            item = NoticeItem()
            item['title'] = new.css('a::attr(title)').extract_first()
            date = new.css('span::text').extract_first()
            if date:
                item['date'] = re.sub(r"\.", "-", date)
            else:
                continue
            item['href'] = response.urljoin(new.css('a::attr("href")').extract_first())
            item['_type'] = "sjtuNotice"
            yield item

    # 爬取生活园区通知
    def parseOurHome(self, response):
        news = response.css('.article center[style]')
        for new in news:
            item = NoticeItem()
            item['title'] = new.css('a::text').extract_first()
            item['date'] = "20" + new.css('.date::text').re('[\d-]+')[0]
            item['href'] = response.urljoin(new.css('a::attr("href")').extract_first())
            item['_type'] = "ourHome"
            yield item
