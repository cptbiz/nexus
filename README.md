# Nexus Network Node

Простая установка и настройка ноды для сети Nexus.

## Установка

### Linux (Ubuntu/Debian)

```bash
# Скачать и установить
curl -sSL https://raw.githubusercontent.com/cptbiz/nexus/main/install.sh | bash

# Или вручную
wget https://github.com/cptbiz/nexus/releases/latest/download/nexus-network-linux
chmod +x nexus-network-linux
sudo mv nexus-network-linux /usr/local/bin/nexus-network
```

### macOS

```bash
# Скачать и установить
curl -sSL https://raw.githubusercontent.com/cptbiz/nexus/main/install.sh | bash

# Или вручную
curl -L -o nexus-network-mac https://github.com/cptbiz/nexus/releases/latest/download/nexus-network-mac
chmod +x nexus-network-mac
sudo mv nexus-network-mac /usr/local/bin/nexus-network
```

## Настройка

1. Создайте директорию конфигурации:
```bash
mkdir -p ~/.nexus
```

2. Скопируйте пример конфигурации:
```bash
cp config.example.json ~/.nexus/config.json
```

3. Отредактируйте файл конфигурации `~/.nexus/config.json`:
```json
{
  "environment": "Production",
  "user_id": "YOUR_USER_ID",
  "wallet_address": "YOUR_WALLET_ADDRESS",
  "node_id": "YOUR_NODE_ID"
}
```

## Использование

### Регистрация пользователя
```bash
nexus-network register-user
```

### Регистрация ноды
```bash
nexus-network register-node
```

### Запуск ноды
```bash
nexus-network start
```

### Выход
```bash
nexus-network logout
```

## Команды

- `start` - Запустить prover
- `register-user` - Зарегистрировать нового пользователя
- `register-node` - Зарегистрировать новую ноду или связать существующую с пользователем
- `logout` - Очистить конфигурацию ноды и выйти
- `help` - Показать справку

## Требования

- Linux (Ubuntu 18.04+, Debian 9+) или macOS 10.15+
- Минимум 4GB RAM
- Стабильное интернет-соединение

## Поддержка

Для получения помощи создайте issue в этом репозитории.
