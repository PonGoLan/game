package instances

type InstancesManager struct {
	games map[string]*Instance

	playerHashToRoom map[string]string
}

var (
	manager *InstancesManager
)

func init() {
	manager = new(InstancesManager)

	manager.games = make(map[string]*Instance)
	manager.playerHashToRoom = make(map[string]string)
}

func Create(roomName string) *Instance {
	instance := CreateInstance()
	manager.games[roomName] = instance

	go instance.Run()

	return manager.games[roomName]
}

func GetInstance(roomName string) *Instance {
	return manager.games[roomName]
}

func GetInstanceWithHash(hash string) *Instance {
	roomGame := manager.playerHashToRoom[hash]
	return GetInstance(roomGame)
}

func LinkHashToRoom(hash, roomName string) {
	manager.playerHashToRoom[hash] = roomName
}

func Get() *InstancesManager {
	return manager
}

func (im *InstancesManager) GetInstances() map[string]*Instance {
	return im.games
}
func (im *InstancesManager) NumberOfInstances() int {
	return len(im.games)
}

func (im *InstancesManager) NumberOfPlayers() int {
	return len(im.playerHashToRoom)
}
