import subprocess
import tkinter as tk
from tkinter import scrolledtext
import threading

TARGET_PATH = r"\\MICHAEL\ctshippingapp\STICKER\PROCESS"

def search_folders_powershell(keyword):
    # PowerShell 只搜索文件夹
    cmd = (
        f'powershell -Command "Get-ChildItem -Path \\"{TARGET_PATH}\\" -Directory -Recurse '
        f'-ErrorAction SilentlyContinue | Where-Object {{$_.Name -like \\"*{keyword}*\\"}} '
        f'| Select-Object -ExpandProperty FullName"'
    )
    try:
        result = subprocess.check_output(cmd, shell=True, text=True)
        matches = result.strip().splitlines() if result.strip() else []
    except subprocess.CalledProcessError:
        matches = []

    result_box.after(0, update_result_box, matches)

def update_result_box(matches):
    result_box.delete(1.0, tk.END)
    if matches:
        result_box.insert(tk.END, f"找到 {len(matches)} 个匹配文件夹:\n\n")
        for m in matches:
            result_box.insert(tk.END, m + "\n")
    else:
        result_box.insert(tk.END, "没有找到匹配文件夹。\n")
    search_btn.config(state=tk.NORMAL)

def on_search():
    keyword = entry.get().strip()
    if not keyword:
        result_box.delete(1.0, tk.END)
        result_box.insert(tk.END, "请输入搜索关键字。\n")
        return
    result_box.delete(1.0, tk.END)
    result_box.insert(tk.END, "正在搜索，请稍候...\n")
    search_btn.config(state=tk.DISABLED)
    threading.Thread(target=search_folders_powershell, args=(keyword,), daemon=True).start()

# UI
root = tk.Tk()
root.title("文件夹搜索工具")
root.geometry("600x400")

frame = tk.Frame(root)
frame.pack(pady=10)

tk.Label(frame, text="请输入搜索关键字:").pack(side=tk.LEFT)
entry = tk.Entry(frame, width=30)
entry.pack(side=tk.LEFT, padx=5)
search_btn = tk.Button(frame, text="搜索", command=on_search)
search_btn.pack(side=tk.LEFT)

result_box = scrolledtext.ScrolledText(root, width=80, height=20)
result_box.pack(padx=10, pady=10)

root.mainloop()
