package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Environment string `json:"environment"`
	UserID      string `json:"user_id"`
	WalletAddress string `json:"wallet_address"`
	NodeID      string `json:"node_id"`
}

type NodeStatus struct {
	Status      string    `json:"status"`
	LastUpdate  time.Time `json:"last_update"`
	NodeID      string    `json:"node_id"`
	Uptime      string    `json:"uptime"`
}

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]
	
	switch command {
	case "start":
		startNode()
	case "register-user":
		registerUser()
	case "register-node":
		registerNode()
	case "logout":
		logout()
	case "help":
		showHelp()
	default:
		fmt.Printf("❌ Неизвестная команда: %s\n", command)
		showHelp()
	}
}

func showHelp() {
	fmt.Println(`
🚀 Nexus Network Node - Справка

Команды:
  start           - Запустить ноду
  register-user   - Зарегистрировать пользователя
  register-node   - Зарегистрировать ноду
  logout          - Выйти и очистить конфигурацию
  help            - Показать эту справку

Примеры:
  nexus-network start
  nexus-network register-user
  nexus-network help
`)
}

func startNode() {
	fmt.Println("🚀 Запуск ноды Nexus...")
	
	config := loadConfig()
	if config == nil {
		fmt.Println("❌ Конфигурация не найдена. Сначала зарегистрируйтесь.")
		return
	}
	
	fmt.Printf("✅ Нода запущена с ID: %s\n", config.NodeID)
	fmt.Printf("👤 Пользователь: %s\n", config.UserID)
	fmt.Printf("💰 Кошелек: %s\n", config.WalletAddress)
	fmt.Printf("🌍 Окружение: %s\n", config.Environment)
	
	// Симуляция работы ноды
	fmt.Println("🔄 Нода работает... Нажмите Ctrl+C для остановки")
	
	// Бесконечный цикл для имитации работы ноды
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("⏰ %s - Нода активна\n", time.Now().Format("15:04:05"))
	}
}

func registerUser() {
	fmt.Println("👤 Регистрация пользователя...")
	
	var userID, walletAddress string
	
	fmt.Print("Введите User ID: ")
	fmt.Scanln(&userID)
	
	fmt.Print("Введите адрес кошелька: ")
	fmt.Scanln(&walletAddress)
	
	config := &Config{
		Environment:   "Production",
		UserID:        userID,
		WalletAddress: walletAddress,
		NodeID:        "",
	}
	
	saveConfig(config)
	fmt.Println("✅ Пользователь зарегистрирован!")
}

func registerNode() {
	fmt.Println("🔧 Регистрация ноды...")
	
	config := loadConfig()
	if config == nil {
		fmt.Println("❌ Сначала зарегистрируйте пользователя")
		return
	}
	
	var nodeID string
	fmt.Print("Введите Node ID: ")
	fmt.Scanln(&nodeID)
	
	config.NodeID = nodeID
	saveConfig(config)
	
	fmt.Printf("✅ Нода зарегистрирована с ID: %s\n", nodeID)
}

func logout() {
	fmt.Println("🚪 Выход из системы...")
	
	configPath := getConfigPath()
	if err := os.Remove(configPath); err != nil {
		fmt.Printf("❌ Ошибка при удалении конфигурации: %v\n", err)
		return
	}
	
	fmt.Println("✅ Конфигурация очищена. Выход выполнен.")
}

func loadConfig() *Config {
	configPath := getConfigPath()
	
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil
	}
	
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil
	}
	
	return &config
}

func saveConfig(config *Config) {
	configPath := getConfigPath()
	configDir := filepath.Dir(configPath)
	
	// Создаем директорию если не существует
	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Fatalf("Ошибка создания директории: %v", err)
	}
	
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("Ошибка сериализации конфигурации: %v", err)
	}
	
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		log.Fatalf("Ошибка записи конфигурации: %v", err)
	}
}

func getConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Ошибка получения домашней директории: %v", err)
	}
	
	return filepath.Join(homeDir, ".nexus", "config.json")
}
