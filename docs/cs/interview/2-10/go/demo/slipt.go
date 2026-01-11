package main(){
split()
}

// input: [1, 3, 2, 6, 7, 2, 1, 13 ,20, 12]  正整数的数组  
// ouput: 左边是奇数、右边是偶数
func split() {
  arr := []int{1, 3, 2, 6, 7, 2, 1, 13 ,20, 12}
  i,j:=0,len(arr-1)
  for i < j {
    if arr[i] % 2 != 0 {
      i++
	  continue
    } 
    if arr[j] % 2 == 0 {
      j--
	  continue
    }
    arr[i], arr[j] = arr[j], arr[i]
    i++
    j--
  }
}