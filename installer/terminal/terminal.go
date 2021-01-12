package terminal

type Terminal struct {
	Stdout bytes.Buffer
	Stderr bytes.Buffer
	

	History []string
}
