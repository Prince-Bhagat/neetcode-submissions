
func topKFrequent(nums []int, k int) []int {
	table := make(map[int]int)
	heap := NewHeap()
	for _,num := range nums{
		_, isExist := table[num]
		if(isExist){
			table[num]++
		}else{
			table[num] = 1
		}
	}
	
	for num, freq := range table{
		heap.addElement(Item{num: num, freq: freq})

	}
	
	result := []int{}
	for i := 0; i < k; i++ {
		
		result = append(result, heap.removeMax().num)
	}
	return result
}



type Item struct {
	num int
	freq int
}


type Heap  struct{
	array []Item
}

func NewHeap()(*Heap){
	return &Heap{
		array : []Item{},
	}
}
func (heap * Heap) addElement(item Item){
	heap.array = append(heap.array,item)
	currentIndex := len(heap.array) -1
	for(currentIndex != 0){

		parentIndex := (currentIndex-1)/2
		
		if(heap.array[parentIndex].freq < heap.array[currentIndex].freq){
			heap.array[parentIndex] , heap.array[currentIndex] = heap.array[currentIndex], heap.array[parentIndex]
		}
		currentIndex = parentIndex
	}

}

func (heap *Heap)removeMax()Item{
	result := heap.array[0]


	heap.array[0] = heap.array[len(heap.array) -1]

	currentIndex := 0
	leftChild , rightChild := 0, 0
	for (currentIndex< len(heap.array)){
		leftChild = (2*currentIndex) +1
		rightChild = (2*currentIndex) +2

		if(leftChild < len(heap.array) && rightChild < len(heap.array) && heap.array[leftChild].freq > heap.array[rightChild].freq  && heap.array[currentIndex].freq < heap.array[leftChild].freq){
			heap.array[currentIndex], heap.array[leftChild] = heap.array[leftChild],heap.array[currentIndex] 
			currentIndex = leftChild
			continue
		}
		if(rightChild < len(heap.array) && heap.array[currentIndex].freq < heap.array[rightChild].freq){
			heap.array[currentIndex], heap.array[rightChild] = heap.array[rightChild],heap.array[currentIndex] 
			currentIndex = rightChild
			continue
		}
		break

	}
	heap.array = heap.array[:len(heap.array) -1]
	return result
}

