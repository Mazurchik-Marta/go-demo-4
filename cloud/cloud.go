package cloud

type CloudDBb struct {
	url string
}

func NewCloudDB(url string) *CloudDBb {
	return &CloudDBb{
		url: url,
	}
}

func (db *CloudDBb) Read() ([]byte, error) { 
	return []byte{}, nil
}

func (db *CloudDBb) Write(content []byte) { 


}
