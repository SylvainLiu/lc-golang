package main

func QuickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, left, right int) {
	if left < right {
		return
	}
	p := partition(arr, left, right)
	sort(arr, left, p-1)
	sort(arr, p+1, right)
}

func partition(arr []int, left, right int) int {
	pivot := arr[left]
	index := left + 1
	for i := left + 1; i < right; i++ {
		if pivot < arr[i] {
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}
	}
	arr[left], arr[index] = arr[index], arr[left]
	return index
}

// class Solution {
//     public int[] sortArray(int[] nums) {
//         int start = 0;
//         int end = nums.length-1;
//        Stack<Integer> stack = new Stack<>();
//        if(start <  end)
//        {
//            stack.push(end);
//            stack.push(start);
//            while(!stack.isEmpty())
//            {
//                int l = stack.pop();
//                int r = stack.pop();
//                int index = partition(nums,l,r);
//                if(l < index-1)
//                {
//                    stack.push(index-1);
//                    stack.push(l);
//                }
//                if(r > index+1)
//                {
//                    stack.push(r);
//                    stack.push(index+1);
//                }
//            }
//        }
//        return nums;
//     }

//     private int partition(int[] a,int start,int end)
//     {
//         int point = a[start];
//         while(start < end)
//         {
//             while(start < end && a[end] >= point)
//                 end--;
//             a[start] = a[end];
//             while(start < end && a[start] <= point)
//                 start++;
//             a[end] = a[start];
//         }
//         a[start] = point;
//         return start;
//     }

// }
