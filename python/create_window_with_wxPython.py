# https://docs.wxpython.org/

import wx

wx_app = wx.App()

# Frameの作成
wx_frame = wx.Frame(parent=None, id=1, title='wx window', size=(600, 400))

wx_frame.Show()
wx_app.SetTopWindow(wx_frame)

wx_app.MainLoop()
