# -*- coding: utf-8 -*-
# defaultencoding = 'utf-8'

import os
import time
import sys
import urllib2

'''
filename   待上传的文件
fieldname  表单域中的name属性
'''
def CreateBody(filename, fieldname, strBoundary):
  bRet = False
  sData = []
  sData.append('--%s' % strBoundary)
  #'Content-Disposition: form-data; name="uploadfile"; filename="XX-Net-1.3.6.zip"'
  sData.append('Content-Disposition: form-data; name="%s";' % fieldname + 'filename="%s"' % os.path.basename(filename))
  sData.append('Content-Type: %s\r\n' % 'application/octet-stream')

  try:
    pFile = open(filename, 'rb')
    sData.append(pFile.read())
    sData.append('--%s--\r\n' % strBoundary)
    bRet = True
  finally:
    pFile.close()

  return bRet, sData

def uploadfile(http_url, filename, fieldname):
  if os.path.exists(filename):
    filesize = os.path.getsize(filename)
    print('file:' + filename + ' is %d bytes!' % filesize)
  else:
    print('file:' + filename + ' isn\'t exists!')
    return False

  strBoundary = '---------------------------%s' % hex(int(time.time() * 1000))
  bRet, sBodyData = CreateBody(filename, fieldname, strBoundary)
  if True == bRet:
    http_body = '\r\n'.join(sBodyData)
    stReq = urllib2.Request(http_url, http_body)
    stReq.add_header('User-Agent','Mozilla/5.0')
    # stReq.add_header('Content-Length:', '%d' % filesize)
    stReq.add_header('Content-Type', 'multipart/form-data; boundary=%s' % strBoundary)
    resp = urllib2.urlopen(stReq, timeout=5)
    # get response
    msg = resp.read()
    print("Response content:\n" + msg)
  else:
    print("CreateBody failed!")

  return bRet

if __name__ == '__main__':

  # if len(sys.argv) > 2:
    # http_url = sys.argv[1]
  # http_url = "http://127.0.0.1:8089/hpnh_upload"
  # http_url = "http://172.20.72.33:8089/hpnh_upload"
  http_url = "http://172.20.72.49:8080/hpnh_upload"
  # filename = "python_upload_test.txt"
  filename = "upload.xml"
  # else:
  print('python upload.py http://127.0.0.1:8089/hpnh_upload ./python_upload_test.txt')
  # sys.exit()

  # 参数3 "uploadfile" 是post表单中的name属性，需要与服务端保持一致
  uploadfile(http_url, filename, 'uploadfile')




    