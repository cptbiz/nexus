#!/bin/bash

set -e

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}🚀 Установка Nexus Network Node...${NC}"

# Определяем ОС
OS="$(uname -s)"
case "${OS}" in
    Linux*)     MACHINE=Linux;;
    Darwin*)    MACHINE=Mac;;
    CYGWIN*)    MACHINE=Cygwin;;
    MINGW*)     MACHINE=MinGw;;
    *)          MACHINE="UNKNOWN:${OS}"
esac

echo -e "${YELLOW}📱 Обнаружена ОС: ${MACHINE}${NC}"

# Создаем директорию
NEXUS_DIR="$HOME/.nexus"
echo -e "${YELLOW}📁 Создаем директорию: ${NEXUS_DIR}${NC}"
mkdir -p "$NEXUS_DIR"

# Определяем архитектуру
ARCH="$(uname -m)"
case "${ARCH}" in
    x86_64)     ARCH="x86_64";;
    aarch64)    ARCH="aarch64";;
    arm64)      ARCH="aarch64";;
    *)          ARCH="x86_64"
esac

echo -e "${YELLOW}🔧 Архитектура: ${ARCH}${NC}"

# Скачиваем бинарный файл
if [ "$MACHINE" = "Linux" ]; then
    BINARY_NAME="nexus-network-linux"
    DOWNLOAD_URL="https://github.com/cptbiz/nexus/releases/latest/download/nexus-network-linux"
elif [ "$MACHINE" = "Mac" ]; then
    BINARY_NAME="nexus-network-mac"
    DOWNLOAD_URL="https://github.com/cptbiz/nexus/releases/latest/download/nexus-network-mac"
else
    echo -e "${RED}❌ Неподдерживаемая ОС: ${OS}${NC}"
    exit 1
fi

echo -e "${YELLOW}⬇️  Скачиваем ${BINARY_NAME}...${NC}"

# Скачиваем файл
if command -v curl >/dev/null 2>&1; then
    curl -L -o "$NEXUS_DIR/$BINARY_NAME" "$DOWNLOAD_URL"
elif command -v wget >/dev/null 2>&1; then
    wget -O "$NEXUS_DIR/$BINARY_NAME" "$DOWNLOAD_URL"
else
    echo -e "${RED}❌ Не найден curl или wget. Установите один из них.${NC}"
    exit 1
fi

# Делаем файл исполняемым
chmod +x "$NEXUS_DIR/$BINARY_NAME"

# Создаем символическую ссылку в /usr/local/bin
if [ "$MACHINE" = "Linux" ]; then
    sudo ln -sf "$NEXUS_DIR/$BINARY_NAME" /usr/local/bin/nexus-network
else
    ln -sf "$NEXUS_DIR/$BINARY_NAME" /usr/local/bin/nexus-network
fi

# Создаем базовую конфигурацию
if [ ! -f "$NEXUS_DIR/config.json" ]; then
    echo -e "${YELLOW}⚙️  Создаем базовую конфигурацию...${NC}"
    cat > "$NEXUS_DIR/config.json" << EOF
{
  "environment": "Production",
  "user_id": "YOUR_USER_ID",
  "wallet_address": "YOUR_WALLET_ADDRESS",
  "node_id": "YOUR_NODE_ID"
}
EOF
fi

echo -e "${GREEN}✅ Установка завершена!${NC}"
echo -e "${YELLOW}📝 Не забудьте настроить config.json в ${NEXUS_DIR}${NC}"
echo -e "${GREEN}🚀 Запустите ноду командой: nexus-network start${NC}"
echo -e "${GREEN}❓ Для справки: nexus-network help${NC}"
