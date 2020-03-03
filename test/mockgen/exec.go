package mockgen

// Controller 这个结构体演示了一种初始化接口的方式
type Controller struct {
	GetSetter
}

// GetThenSet 检查值是否已设置。如果没有设置就将其设置
func (c *Controller) GetThenSet(key, value string) error {
	val, err := c.Get(key)
	if err != nil {
		return err
	}

	if val != value {
		return c.Set(key, value)
	}
	return nil
}
