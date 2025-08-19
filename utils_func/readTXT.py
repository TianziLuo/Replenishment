import sys
from pathlib import Path

# print(sys.argv)  
def readTXT(folder: Path, file_path: str):
    target_file = folder / file_path

    if target_file.exists():
        with open(target_file, 'r', encoding='utf-8') as f:
            outputData = f.readline().strip()
        print(f"Value from {file_path}: {outputData}")
        return outputData
    else:
        print(f"{file_path} not found in {folder}")
        return None

def readSKU(folder: Path):
    return readTXT(folder, file_path="sku.txt")

def readOrderNO(folder: Path):
    return readTXT(folder, file_path="orderno.txt")


'''
if __name__ == "__main__":
    sku_value = readSKU()
    order_value= readOrderNO()
    print(f"SKU read from file: {sku_value}")
'''



 # folder = Path(sys.argv[1])