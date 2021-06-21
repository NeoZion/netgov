
import win32api
import win32gui
import time
import win32con

from songs.yhyl import yhyl 

name = "yhyl"

windowName = u"天涯明月刀"

key_map = {
    "1": 49, "2": 50, "3": 51, "4": 52, "5": 53, "6": 54, "7": 55, "8": 56, "9": 57, "0": 58,
    "A": 65, "B": 66, "C": 67, "D": 68, "E": 69, "F": 70, "G": 71, "H": 72, "I": 73, "J": 74,
    "K": 75, "L": 76, "M": 77, "N": 78, "O": 79, "P": 80, "Q": 81, "R": 82, "S": 83, "T": 84,
    "U": 85, "V": 86, "W": 87, "X": 88, "Y": 89, "Z": 90
}

def mouse_click(x):
    win32api.SetCursorPos([x[0]+10,x[1]+10])
    win32api.mouse_event(win32con.MOUSEEVENTF_LEFTDOWN, 0, 0, 0, 0)
    win32api.mouse_event(win32con.MOUSEEVENTF_LEFTUP, 0, 0, 0, 0)
    return True

def getHwnd():
    hwnd = win32gui.FindWindow(0,windowName)
    if (hwnd):
        rect = win32gui.GetWindowRect(hwnd)
        return rect[0],rect[1]
    return None
 
 
def key_down(key):
    key = key.upper()
    vk_code = key_map[key]
    win32api.keybd_event(vk_code,win32api.MapVirtualKey(vk_code,0),0,0)
 
 
def key_up(key):
    key = key.upper()
    vk_code = key_map[key]
    win32api.keybd_event(vk_code, win32api.MapVirtualKey(vk_code, 0), win32con.KEYEVENTF_KEYUP, 0)
 
 
def key_press(key,timeout):
    key_down(key)
    time.sleep(timeout)
    key_up(key)

def play_music(name):
    bpms = 60 / name.source["bpm"]
    for i in name.chapterA:
        for j in i:
            if j[0] == '-':
                time.sleep(bpms*name.source["beat_m"]/j[1])
                continue
            else:
                key_press(j[0],j[1]*name.source["beat_m"]/j[1])



if __name__ == '__main__':
    mouse_click(getHwnd())
    time.sleep(3)
    play_music()
