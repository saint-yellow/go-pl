package business

func init() {
	students = []Student{
		{
			ID: 1,
			Name: "Nick Carter",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type: GradeQuiz,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type: GradeExam,
					Score: 94,
				},
				{
					Title: "Quiz 2",
					Type: GradeQuiz,
					Score: 82,
				},
			},
		},
		{
			ID: 2,
			Name: "Roberto Baggio",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type: GradeQuiz,
					Score: 100,
				},
				{
					Title: "Final Exam",
					Type: GradeExam,
					Score: 100,
				},
				{
					Title: "Quiz 2",
					Type: GradeQuiz,
					Score: 81,
				},
			},
		},
		{
			ID: 3,
			Name: "Emma Stone",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type: GradeQuiz,
					Score: 67,
				},
				{
					Title: "Final Exam",
					Type: GradeExam,
					Score: 0,
				},
				{
					Title: "Quiz 2",
					Type: GradeQuiz,
					Score: 75,
				},
			},
		},
		{
			ID: 4,
			Name: "Rachel McAdams",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type: GradeQuiz,
					Score: 98,
				},
				{
					Title: "Final Exam",
					Type: GradeExam,
					Score: 99,
				},
				{
					Title: "Quiz 2",
					Type: GradeQuiz,
					Score: 94,
				},
			},
		},
		{
			ID: 5,
			Name: "Kelly Clarkson",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type: GradeQuiz,
					Score: 95,
				},
				{
					Title: "Final Exam",
					Type: GradeExam,
					Score: 100,
				},
				{
					Title: "Quiz 2",
					Type: GradeQuiz,
					Score: 97,
				},
			},
		},
	}
}