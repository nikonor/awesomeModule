package awesomeModule

import "unicode"

func AwesomeString(in string) string {
    out := make([]rune,len(in))
    for i,ch := range in {
        if i % 2 == 0 {
            out[i] =  unicode.ToUpper(ch)
        } else {
            out[i] =  unicode.ToLower(ch)
        }
    }
    
    return string(out)
}
