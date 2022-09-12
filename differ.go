package go_differ

func NewDiffer(v any, opts ...Option) Differ {
	o := &options{
		ignoreFields: nil,
	}
	for _, fn := range opts {
		fn(o)
	}

	d := &differ{
		options: o,
	}
	d.snapshot = d.getSnapshot(v)
	return d
}

type Differ interface {
	IsDirty(v any) bool
	GetChangeKeys(v any) []string
	GetChanges(v any) map[string]any
}

type differ struct {
	*options
	// 快照
	snapshot EntryValues
}

func (d *differ) IsDirty(v interface{}) (isDirty bool) {
	for key, value := range d.getSnapshot(v) {
		originValue, ok := d.snapshot[key]
		if !ok || originValue != value {
			return true
		}
	}
	return false
}

func (d *differ) GetChangeKeys(v interface{}) (keys []string) {
	for key, value := range d.getSnapshot(v) {
		originValue, ok := d.snapshot[key]
		if !ok || originValue != value {
			keys = append(keys, key)
			continue
		}
	}
	return keys
}

func (d *differ) GetChanges(v interface{}) (changes map[string]interface{}) {
	changes = make(map[string]interface{}, 0)
	for key, value := range d.getSnapshot(v) {
		originValue, ok := d.snapshot[key]
		if !ok || originValue != value {
			changes[key] = value
			continue
		}
	}
	return changes
}
