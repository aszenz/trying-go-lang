package iteration

// RepeatCharacter take a character and repeats it by repetition no of times
func RepeatCharacter(character string, repetition int) string {
	var repeated string
	for i:=0; i < repetition; i++ {
		repeated += character 
	}
	return repeated 
}