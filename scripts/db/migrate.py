import os
import subprocess

# 配置参数
DB_NAME = "erp_system"
DB_USER = "postgres"
DB_PASSWORD = "77072166"
DB_HOST = "localhost"
DB_PORT = 5432
MIGRATIONS_PATH = r"D:\ERP_SYSTEM\backend\migrations"

def run_migrations():
    # 构建数据库URL
    db_url = f"postgres://{DB_USER}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_NAME}?sslmode=disable"
    
    # 构建迁移命令
    cmd = f"migrate -path {MIGRATIONS_PATH} -database '{db_url}' up"
    
    try:
        subprocess.run(cmd, shell=True, check=True)
        print("数据库迁移成功！")
    except subprocess.CalledProcessError as e:
        print(f"迁移失败: {e}")
        exit(1)

if __name__ == "__main__":
    run_migrations() 