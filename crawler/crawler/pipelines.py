# -*- coding: utf-8 -*-
import os
import sqlite3


# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


class CrawlerPipeline:
    def process_item(self, item, spider):
        return item


# NoticePipeline 进行连接数据库等初始化操作，并将符合条件的item存入数据库
class NoticePipeline:
    def __init__(self):
        self.conn = sqlite3.connect(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'bulletin.db'))
        self.cursor = self.conn.cursor()

    def process_item(self, item, spider):
        # 将符合条件的item存入数据库，并且之存入最新通知
        if item['title'] and item['href'] and item['date']:
            try:
                self.cursor.execute(
                    """INSERT INTO notices (title, href, `date`, `type`) VALUES ('%s', '%s', '%s', '%s')""" % (
                        item['title'], item['href'], item['date'], item['_type']))
                self.conn.commit()  # commit是必须的
            except:
                print("Duplicated notice")
            return item

    def spider_closed(self, spider):
        self.conn.close()  # close database
