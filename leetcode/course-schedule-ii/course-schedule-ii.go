package course_schedule_ii

func include(list []int, n int) bool {
	for _, v := range list {
		if v == n {
			return true
		}
	}
	return false
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	// prerequisites[i].length = 2
	// first, find roots(never appear as first element), then add roots to queue
	restDependentElements := make(map[int]bool, 0)
	for _, e := range prerequisites {
		restDependentElements[e[0]] = true
	}
	courseQueue := []int{}
	for i := 0; i < numCourses; i++ {
		if !restDependentElements[i] {
			courseQueue = append(courseQueue, i)
		}
	}
	// layer and layer -> int[], int[], int[] ...
	// find from restDependentElements
	// like width first

	// merge depend info course -> depended courses: [][]int
	course2dependedCourse := [][]int{}
	for i := 0; i < numCourses; i++ {
		course2dependedCourse = append(course2dependedCourse, []int{})
	}
	for _, vector := range prerequisites {
		c, dc := vector[0], vector[1]
		course2dependedCourse[c] = append(course2dependedCourse[c], dc)
	}

	for len(courseQueue) != numCourses {
		currentLen := len(courseQueue)
		// find courses who depend on courses contained in courseQueue
		for c := 0; c < len(course2dependedCourse); c++ {
			if include(courseQueue, c) {
				continue
			}
			// TODO: if dependedCourse all in courseQueue, add course to courseQueue
			dependedCourse := course2dependedCourse[c]
			allIn := true
			for _, dc := range dependedCourse {
				if !include(courseQueue, dc) {
					allIn = false
					break
				}
			}
			if allIn {
				courseQueue = append(courseQueue, c)
			}
		}
		if currentLen == len(courseQueue) {
			return []int{}
		}
	}

	// TODO: extreme conditions

	// result
	return courseQueue
}
