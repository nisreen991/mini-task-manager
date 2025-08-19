import { Routes } from '@angular/router';
import { AuthFormComponent } from './auth-form/auth-form';
import { TasksListComponent } from './tasks-list/tasks-list';
import { AddEditTaskComponent } from './add-edit-task/add-edit-task';
import { AuthGuard } from './auth.guard'; // We will create this file next

export const routes: Routes = [
  { path: 'login', component: AuthFormComponent },
  { path: 'signup', component: AuthFormComponent },
  { path: '', redirectTo: '/login', pathMatch: 'full' },
  { path: 'tasks', component: TasksListComponent, canActivate: [AuthGuard] },
  { path: 'add-edit-task', component: AddEditTaskComponent, canActivate: [AuthGuard] },
  { path: 'add-edit-task/:id', component: AddEditTaskComponent, canActivate: [AuthGuard] },
  { path: '**', redirectTo: '/login' }
];
