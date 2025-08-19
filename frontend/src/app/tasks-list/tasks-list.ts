import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TaskService } from '../task';
import { Router } from '@angular/router';

interface Task {
  id?: string;
  title: string;
  description: string;
  completed: boolean;
}

@Component({
  selector: 'app-tasks-list',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './tasks-list.html',
  styleUrl: './tasks-list.scss',
})
export class TasksListComponent implements OnInit {
  tasks: Task[] = [];

  constructor(private taskService: TaskService, private router: Router) {}

  ngOnInit(): void {
    this.fetchTasks();
  }

  fetchTasks(): void {
    this.taskService.getTasks().subscribe(
      (tasks: Task[]) => {
        this.tasks = tasks;
      },
      (error: any) => {
        console.error('Error fetching tasks:', error);
        // Handle error, e.g., redirect to login if unauthorized
        if (error.status === 401) {
          this.router.navigate(['/login']);
        }
      }
    );
  }

  toggleComplete(task: Task): void {
    task.completed = !task.completed;
    this.taskService.updateTask(task.id!, task).subscribe(
      () => {
        console.log('Task updated successfully');
      },
      (error: any) => {
        console.error('Error updating task:', error);
        // Revert UI change if update fails
        task.completed = !task.completed;
      }
    );
  }

  editTask(id: string): void {
    this.router.navigate(['/add-edit-task', id]);
  }

  deleteTask(id: string): void {
    if (confirm('Are you sure you want to delete this task?')) {
      this.taskService.deleteTask(id).subscribe(
        () => {
          this.tasks = this.tasks.filter((task) => task.id !== id);
        },
        (error: any) => {
          console.error('Error deleting task:', error);
        }
      );
    }
  }

  addNewTask(): void {
    this.router.navigate(['/add-edit-task']);
  }
}
