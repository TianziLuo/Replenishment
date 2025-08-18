from readSKU import readSKU
from pathlib import Path
from openpyxl import load_workbook

def write2xlax(SKU_line):
# 打开 Excel 文件
    excel_path = Path(r"C:\Frank\补发货.xlsx")
    wb = load_workbook(excel_path)

    # 获取 "InHand" 表
    ws = wb["InHand"]

    # 找到最后一行
    last_row = ws.max_row

    # 写入最后一行下一行的 A 列
    ws.cell(row=last_row + 1, column=1, value=SKU_line)

    # 保存 Excel 文件
    wb.save(excel_path)
