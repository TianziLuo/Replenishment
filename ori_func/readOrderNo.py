import sys
from pathlib import Path

def readSKU():
    # folder path from GO
    folder = Path(sys.argv[1])
    sku_file = folder / "orderno.txt"

    if sku_file.exists():
        # read A1
        with open(sku_file, 'r', encoding='utf-8') as f:
            orderNo_line = f.readline().strip()
        print(f"SKU value: {orderNo_line}")

    else:
        print("sku.TXT not found in the folder")
    return orderNo_line