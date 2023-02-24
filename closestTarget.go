package main

// circular string
// find nearest word
// words[(i + 1) % n] --> next word of words[i]
// words[(i - 1 + n) % n] --> prev word of words[i]

func closetTarget(words []string, target string, startIndex int) int {
	var (
		prevIndex  int
		nextIndex  int
		i          int
		rigthCount int
		leftCount  int
	)
	if words[startIndex] == target {
		return 0
	}
	for i = 0; i < len(words); i++ {

		if words[nextIndex] == target {
			return rigthCount
		}
		if words[prevIndex] == target {
			return leftCount
		}
		leftCount++
		rigthCount++
		nextIndex = (nextIndex + 1) % len(words)
		prevIndex = (prevIndex - 1 + len(words)) % len(words)
	}
	return -1
}

// class Solution {
//     public int closetTarget(String[] words, String target, int startIndex) {
//         int left = startIndex, right = startIndex, leftCount = 0, rightCount = 0, n = words.length;
//         for (int i = 0; i < n; i++) {
//             if (target.equals(words[left]))
//                 return leftCount;
//             if (target.equals(words[right]))
//                 return rightCount;
//             left = (left - 1 + n) % n;
//             right = (right + 1) % n;
//             leftCount++;
//             rightCount++;
//         }

//         return -1;
//     }
// }
