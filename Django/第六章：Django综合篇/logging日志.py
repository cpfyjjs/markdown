# _*_coding:utf-8_*_


import logging
import logging.config

LOGGING = {
    'version':1,
    'disable_existing_loggers':False,
    'handlers':{
        'console':{
            'level':'DEBUG',
            'class':'logging.StreamHandler',
        },
        'file':{
            'level':'DEBUG',
            'class':'logging.FileHandler',
            'filename':'D:/markdown/Django/第六章：Django综合篇/debug.log',
            'encoding':'utf-8',
        },
    },
    'loggers':{
        'django':{
            'handlers':['console','file'],
            'level':'DEBUG',      
        },
    },
}

# 读取配置信息
logging.config.dictConfig(LOGGING)

# 获取logger记录器
logger = logging.getLogger('django') 
# 使用日志功能
logger.debug('debug message')
logger.info('info message')
logger.warn('warn message')
logger.error('error message')
logger.critical('critical message')




# import logging

# # 创建一个logging记录器
# logger = logging.getLogger('simple_logger')

# # 设置等级
# logger.setLevel(logging.DEBUG)

# # 创建一个控制台处理器，并将日志级别设置为debug
# ch = logging.StreamHandler()
# ch.setLevel(logging.DEBUG)

# # 创建formatter格式化器
# formatter = logging.Formatter('%(asctime)s - %(name)s -%(levelname)s - %(message)s')

# # 将formatter添加到ch处理器
# ch.setFormatter(formatter)

# # 将ch添加到logger
# logger.addHandler(ch)

# # 将logging输出到文件
# fh = logging.FileHandler('D:/markdown/Django/第六章：Django综合篇/debug.log',encoding='utf-8')
# fh.setLevel(logging.DEBUG)
# fh.setFormatter(formatter)
# logger.addHandler(fh)


# logger.debug('debug message')
# logger.info('info message')
# logger.warn('warn message')
# logger.error('error message')
# logger.critical('critical message')