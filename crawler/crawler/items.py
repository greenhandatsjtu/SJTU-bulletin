# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

import scrapy


class CrawlerItem(scrapy.Item):
    # define the fields for your item here like:
    # name = scrapy.Field()
    pass


# 定义了消息的结构
class NoticeItem(scrapy.Item):
    date = scrapy.Field()  # 发布日期
    title = scrapy.Field()  # 标题
    href = scrapy.Field()  # 超链接
    _type = scrapy.Field()  # 类型
