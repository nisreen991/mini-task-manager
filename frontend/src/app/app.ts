import { Component } from '@angular/core';
import { RouterOutlet, Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { AuthService } from './auth';
import { LoadingSpinnerComponent } from './loading-spinner/loading-spinner';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet, LoadingSpinnerComponent],
  template: `
    <header *ngIf="authService.isAuthenticated()">
      <h1>Mini Task Manager</h1>
      <button (click)="logout()">Logout</button>
    </header>
    <main>
      <router-outlet></router-outlet>
    </main>
    <app-loading-spinner></app-loading-spinner>
  `,
})
export class AppComponent {
  title = 'task-manager-app';

  constructor(public authService: AuthService, private router: Router) {}

  logout(): void {
    this.authService.logout();
    this.router.navigate(['/login']);
  }
}
