import sys
from pathlib import Path

def readSKU():
    # folder path from GO
    folder = Path(sys.argv[1])
    sku_file = folder / "sku.txt"

    if sku_file.exists():
        # read A1
        with open(sku_file, 'r', encoding='utf-8') as f:
            SKU_line = f.readline().strip()
        print(f"SKU value: {SKU_line}")

    else:
        print("sku.TXT not found in the folder")
    return SKU_line