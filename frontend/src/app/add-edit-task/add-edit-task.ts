import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { TaskService } from '../task';

interface Task {
  id?: string;
  title: string;
  description: string;
  completed: boolean;
}

@Component({
  selector: 'app-add-edit-task',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './add-edit-task.html',
  styleUrl: './add-edit-task.scss',
})
export class AddEditTaskComponent implements OnInit {
  task: Task = { title: '', description: '', completed: false };
  isEditMode = false;
  taskId: string | null = null;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private taskService: TaskService
  ) {}

  ngOnInit(): void {
    this.taskId = this.route.snapshot.paramMap.get('id');
    if (this.taskId) {
      this.isEditMode = true;
      this.taskService.getTask(this.taskId).subscribe(
        (task: Task) => {
          this.task = task;
        },
        (error: any) => {
          console.error('Error fetching task:', error);
          // Handle error, e.g., navigate back to tasks list
          this.router.navigate(['/tasks']);
        }
      );
    }
  }

  onSubmit(): void {
    if (this.isEditMode) {
      this.taskService.updateTask(this.taskId!, this.task).subscribe(
        () => {
          this.router.navigate(['/tasks']);
        },
        (error: any) => {
          console.error('Error updating task:', error);
        }
      );
    } else {
      this.taskService.createTask(this.task).subscribe(
        () => {
          this.router.navigate(['/tasks']);
        },
        (error: any) => {
          console.error('Error creating task:', error);
        }
      );
    }
  }

  cancel(): void {
    this.router.navigate(['/tasks']);
  }
}
