// package email

// import (
// 	"fmt"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/sriniously/gotodo/internal/model/todo"
// )

// func (c *Client) SendWelcomeEmail(to, firstName string) error {
// 	data := map[string]any{
// 		"UserFirstName": firstName,
// 	}

// 	return c.SendEmail(
// 		to,
// 		"Welcome to Boilerplate!",
// 		TemplateWelcome,
// 		data,
// 	)
// }

// func (c *Client) SendDueDateReminderEmail(to, todoTitle string, todoID uuid.UUID, dueDate time.Time) error {
// 	data := map[string]interface{}{
// 		"TodoTitle":    todoTitle,
// 		"TodoID":       todoID.String(),
// 		"DueDate":      dueDate.Format("Monday, January 2, 2006 at 3:04 PM"),
// 		"DaysUntilDue": int(dueDate.Sub(time.Now()).Hours() / 24),
// 	}

// 	return c.SendEmail(
// 		to,
// 		fmt.Sprintf("Reminder: '%s' is due soon", todoTitle),
// 		TemplateDueDateReminder,
// 		data,
// 	)
// }

// func (c *Client) SendOverdueNotificationEmail(to, todoTitle string, todoID uuid.UUID, dueDate time.Time) error {
// 	data := map[string]interface{}{
// 		"TodoTitle":   todoTitle,
// 		"TodoID":      todoID.String(),
// 		"DueDate":     dueDate.Format("Monday, January 2, 2006 at 3:04 PM"),
// 		"DaysOverdue": int(time.Now().Sub(dueDate).Hours() / 24),
// 	}

// 	return c.SendEmail(
// 		to,
// 		fmt.Sprintf("Overdue: '%s' needs your attention", todoTitle),
// 		TemplateOverdueNotification,
// 		data,
// 	)
// }

// func (c *Client) SendWeeklyReportEmail(to string, weekStart, weekEnd time.Time,
// 	completedCount, activeCount, overdueCount int, completedTodos, overdueTodos []todo.PopulatedTodo,
// ) error {
// 	data := map[string]interface{}{
// 		"WeekStart":      weekStart.Format("January 2, 2006"),
// 		"WeekEnd":        weekEnd.Format("January 2, 2006"),
// 		"CompletedCount": completedCount,
// 		"ActiveCount":    activeCount,
// 		"OverdueCount":   overdueCount,
// 		"CompletedTodos": completedTodos,
// 		"OverdueTodos":   overdueTodos,
// 		"HasCompleted":   completedCount > 0,
// 		"HasOverdue":     overdueCount > 0,
// 	}

// 	return c.SendEmail(
// 		to,
// 		fmt.Sprintf("Your Weekly Productivity Report (%s - %s)",
// 			weekStart.Format("Jan 2"), weekEnd.Format("Jan 2")),
// 		TemplateWeeklyReport,
// 		data,
// 	)
// }


package email

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sriniously/gotodo/internal/model/todo"
)

// --------------------
// Welcome Email
// --------------------

func (c *Client) SendWelcomeEmail(to, firstName string) error {
	data := map[string]any{
		"UserFirstName": firstName,
	}

	return c.SendEmail(
		to,
		"Welcome to Boilerplate!",
		TemplateWelcome,
		data,
	)
}

// --------------------
// Due Date Reminder Email
// --------------------

func (c *Client) SendDueDateReminderEmail(to, todoTitle string, todoID uuid.UUID, dueDate time.Time) error {

	daysUntil := int(dueDate.Sub(time.Now()).Hours() / 24)

	data := map[string]any{
		"TodoTitle":    todoTitle,
		"TodoID":       todoID.String(),
		"DueDate":      dueDate.Format("Monday, January 2, 2006 at 3:04 PM"),
		"DaysUntilDue": daysUntil,
	}

	return c.SendEmail(
		to,
		fmt.Sprintf("Reminder: '%s' is due soon", todoTitle),
		TemplateDueDateReminder,
		data,
	)
}

// --------------------
// Overdue Email
// --------------------

func (c *Client) SendOverdueNotificationEmail(to, todoTitle string, todoID uuid.UUID, dueDate time.Time) error {

	daysOverdue := int(time.Now().Sub(dueDate).Hours() / 24)

	data := map[string]any{
		"TodoTitle":   todoTitle,
		"TodoID":      todoID.String(),
		"DueDate":     dueDate.Format("Monday, January 2, 2006 at 3:04 PM"),
		"DaysOverdue": daysOverdue,
	}

	return c.SendEmail(
		to,
		fmt.Sprintf("Overdue: '%s' needs your attention", todoTitle),
		TemplateOverdueNotification,
		data,
	)
}

// --------------------
// Weekly Report Email
// --------------------

func (c *Client) SendWeeklyReportEmail(
	to string,
	weekStart, weekEnd time.Time,
	completedCount, activeCount, overdueCount int,
	completedTodos, overdueTodos []todo.PopulatedTodo,
) error {

	data := map[string]any{
		"WeekStart":      weekStart.Format("January 2, 2006"),
		"WeekEnd":        weekEnd.Format("January 2, 2006"),
		"CompletedCount": completedCount,
		"ActiveCount":    activeCount,
		"OverdueCount":   overdueCount,
		"CompletedTodos": completedTodos,
		"OverdueTodos":   overdueTodos,
		"HasCompleted":   completedCount > 0,
		"HasOverdue":     overdueCount > 0,
	}

	return c.SendEmail(
		to,
		fmt.Sprintf(
			"Your Weekly Productivity Report (%s - %s)",
			weekStart.Format("Jan 2"),
			weekEnd.Format("Jan 2"),
		),
		TemplateWeeklyReport,
		data,
	)
}