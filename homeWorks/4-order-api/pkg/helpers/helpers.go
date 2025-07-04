package helpers

import "strconv"

func StingToUint(s string)(uint, error){
u, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return  0,err
		}
return uint(u),nil
}

