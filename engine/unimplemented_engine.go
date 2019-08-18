package engine

import "fmt"

type UnimplementedEngine struct{}

func (engine *UnimplementedEngine) Get(key string) (*string, error) {
	return nil, fmt.Errorf("unimplemented method: Get")
}

func (engine *UnimplementedEngine) Set(key string, value string) error {
	return fmt.Errorf("unimplemented method: Set")
}

func (engine *UnimplementedEngine) Del(keys []string) (int64, error) {
	return int64(0), fmt.Errorf("unimplemented method: Del")
}

func (engine *UnimplementedEngine) Exists(keys []string) (int64, error) {
	return int64(0), fmt.Errorf("unimplemented method: Exists")
}

func (engine *UnimplementedEngine) Incr(key string) (int64, error) {
	return int64(0), fmt.Errorf("unimplemented method: Incr")
}

func (engine *UnimplementedEngine) Decr(key string) (int64, error) {
	return int64(0), fmt.Errorf("unimplemented method: Decr")
}

func (engine *UnimplementedEngine) Incrby(key string, by int64) (int64, error) {
	return int64(0), fmt.Errorf("unimplemented method: Incrby")
}

func (engine *UnimplementedEngine) Decrby(key string, by int64) (int64, error) {
	return int64(0), fmt.Errorf("unimplemented method: Decrby")
}

func (engine *UnimplementedEngine) Strlen(key string) (int64, error) {
	return int64(0), fmt.Errorf("unimplemented method: Strlen")
}

func (engine *UnimplementedEngine) GetSet(key string, value string) (*string, error) {
	return nil, fmt.Errorf("unimplemented method: GetSet")
}

func (engine *UnimplementedEngine) MGet(keys []string) ([]*string, error) {
	return []*string{}, fmt.Errorf("unimplemented method: MGet")
}

func (engine *UnimplementedEngine) MSet(kvs map[string]string) error {
	return fmt.Errorf("unimplemented method: MSet")
}

func (engine *UnimplementedEngine) Type(key string) (string, error) {
	return "", fmt.Errorf("unimplemented method: Type")
}

func (engine *UnimplementedEngine) Expire(key string, seconds int64) (bool, error) {
	return false, fmt.Errorf("unimplemented method: Expire")
}

func (engine *UnimplementedEngine) PExpire(key string, millis int64) (bool, error) {
	return false, fmt.Errorf("unimplemented method: PExpire")
}

func (engine *UnimplementedEngine) ExpireAt(key string, seconds int64) (bool, error) {
	return false, fmt.Errorf("unimplemented method: ExpireAt")
}

func (engine *UnimplementedEngine) PExpireAt(key string, millis int64) (bool, error) {
	return false, fmt.Errorf("unipmelemnted method: PExpireAt")
}

func (engine *UnimplementedEngine) Persist(key string) (bool, error) {
	return false, fmt.Errorf("unimplemented method: Persist")
}

func (engine *UnimplementedEngine) TTL(key string) (int64, error) {
	return int64(0), fmt.Errorf("unimplemented method: TTL")
}

func (engine *UnimplementedEngine) PTTL(key string) (int64, error) {
	return int64(0), fmt.Errorf("unimplemented method: PTTL")
}
