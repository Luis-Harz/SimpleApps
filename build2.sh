#!/usr/bin/env python3
import os
import subprocess
import shutil
from zipfile import ZipFile
import paramiko
import getpass

subprocess.run(["rm", "-rf", "build/*"])

# ==== CONFIG ====
PROJECT_ROOT = os.path.dirname(os.path.abspath(__file__))  # Skript im Projekt-Root
MAIN_FILE = os.path.join(PROJECT_ROOT, "main.go")
BUILD_DIR = os.path.join(PROJECT_ROOT, "build")
FOR_SERVER_DIR = os.path.join(BUILD_DIR, "ForServer")
CMD_PATH = os.path.join(BUILD_DIR, "cmd")
# ==== SFTP CONFIG ====
SFTP_HOST = "100.101.119.31"
SFTP_USER = "luis"
SFTP_REMOTE_DIR = "/home/luis/SimpleApps"
SFTP_PORT = 22

# ==== Version eingeben ====
version = input("Gib die Version ein: ").strip()
version_txt = f"{version}\n"
#Tools anzahl angeben
tools = input("Gib die Tools anzahl an: ").strip()
tools_txt = f"{tools}\n"

# ==== Hilfsfunktion ZIP erstellen ====
def make_zip(zip_name, folder):
    folder = os.path.abspath(folder)
    base_folder = os.path.basename(folder)
    with ZipFile(zip_name, 'w') as zipf:
        for root, dirs, files in os.walk(folder):
            dirs[:] = [d for d in dirs if not d.startswith('.')]
            visible_files = [f for f in files if not f.startswith('.')]
            for file in visible_files:
                full_path = os.path.join(root, file)
                rel_path = os.path.relpath(full_path, os.path.dirname(folder))
                zipf.write(full_path, rel_path)
    print(f"{zip_name} erstellt!")


# ==== Build-Ordner erstellen ====
os.makedirs(BUILD_DIR, exist_ok=True)
os.makedirs(FOR_SERVER_DIR, exist_ok=True)

# ==== Version.txt erstellen ====
version_file_path = os.path.join(BUILD_DIR, "version.txt")
with open(version_file_path, "w") as f:
    f.write(version_txt)
#tools.txt erstellen
tools_file_path = os.path.join(BUILD_DIR, "tools.txt")
with open(tools_file_path, "w") as f:
    f.write(tools_txt)

# ==== Go-Modul prüfen ====
if not os.path.exists(os.path.join(PROJECT_ROOT, "go.mod")):
    subprocess.run(["go", "mod", "init", "SimpleApps"], check=True)

# ==== Kompiliere Linux ====
linux_build_dir = os.path.join(BUILD_DIR, "Linux")
os.makedirs(linux_build_dir, exist_ok=True)
linux_output = os.path.join(linux_build_dir, "SimpleApps")
subprocess.run(["go", "build", "-o", linux_output, MAIN_FILE], check=True)
shutil.copy(version_file_path, linux_build_dir)
make_zip(os.path.join(BUILD_DIR, "Linux.zip"), linux_build_dir)

# ==== Kompiliere Windows ====
windows_build_dir = os.path.join(BUILD_DIR, "Windows")
os.makedirs(windows_build_dir, exist_ok=True)
windows_output = os.path.join(windows_build_dir, "SimpleApps.exe")
subprocess.run(["go", "build", "-o", windows_output, MAIN_FILE],
               env={**os.environ, "GOOS":"windows","GOARCH":"amd64"}, check=True)
shutil.copy(version_file_path, windows_build_dir)
make_zip(os.path.join(BUILD_DIR, "Windows.zip"), windows_build_dir)

# ==== ForServer erstellen ====
shutil.copy(linux_output, os.path.join(FOR_SERVER_DIR, "SimpleApps"))
shutil.copy(windows_output, os.path.join(FOR_SERVER_DIR, "SimpleApps.exe"))
shutil.copy(version_file_path, FOR_SERVER_DIR)
shutil.copy(tools_file_path, FOR_SERVER_DIR)
shutil.copy(MAIN_FILE, CMD_PATH)
# ==== SFTP Upload ====
print("Verbinde zu SFTP und lade Dateien hoch...")
password = getpass.getpass("SFTP Passwort: ")
transport = paramiko.Transport((SFTP_HOST, SFTP_PORT))
transport.connect(username=SFTP_USER, password=password)
sftp = paramiko.SFTPClient.from_transport(transport)

# Remote-Ordner erstellen
def mkdir_p(remote_path):
    try:
        sftp.chdir(remote_path)
    except IOError:
        sftp.mkdir(remote_path)
        sftp.chdir(remote_path)

# Hauptordner
mkdir_p(SFTP_REMOTE_DIR)

# Hilfsfunktion: Upload + überschreiben
def upload_file(local_path, remote_dir, remote_name=None):
    if remote_name is None:
        remote_name = os.path.basename(local_path)
    remote_path = os.path.join(remote_dir, remote_name)
    try:
        sftp.remove(remote_path)
    except IOError:
        pass
    sftp.put(local_path, remote_path)
    print(f"Hochgeladen: {remote_path}")

# Upload ZIPs in Hauptordner
upload_file(os.path.join(BUILD_DIR, "Linux.zip"), SFTP_REMOTE_DIR)
upload_file(os.path.join(BUILD_DIR, "Windows.zip"), SFTP_REMOTE_DIR)

# ForServer-Ordner auf Server
remote_forserver = os.path.join(SFTP_REMOTE_DIR, "ForServer")
mkdir_p(remote_forserver)
for file_name in ["SimpleApps", "SimpleApps.exe", "version.txt", "tools.txt"]:
    upload_file(os.path.join(FOR_SERVER_DIR, file_name), remote_forserver)

#Source Code Upload
make_zip(os.path.join(PROJECT_ROOT, "Source.zip"), PROJECT_ROOT)
upload_file(os.path.join(PROJECT_ROOT, "Source.zip"), remote_forserver)

sftp.close()
transport.close()
print("Leere Cache...")
result = subprocess.run(["bash", "cache.sh"], capture_output=True, text=True)
print("STDOUT:", result.stdout)
print("STDERR:", result.stderr)
print("Alle Dateien erfolgreich hochgeladen!")
