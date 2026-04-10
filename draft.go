package draft

type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

//w.Radius=5
//这样可以直接用，就是一种匿名
