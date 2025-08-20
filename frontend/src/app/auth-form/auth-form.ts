import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { AuthService } from '../auth';
import { Router } from '@angular/router';

@Component({
  selector: 'app-auth-form',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './auth-form.html',
  styleUrl: './auth-form.scss',
})
export class AuthFormComponent {
  user = { username: '', password: '' };
  isLoginMode = true;
  errorMessage: string | null = null;
  successMessage: string | null = null;

  constructor(private authService: AuthService, private router: Router) {}

  onSwitchMode() {
    this.isLoginMode = !this.isLoginMode;
    this.errorMessage = null; // Clear error when switching modes
    this.user.username = '';
    this.user.password = ''
  }

  onSubmit() {
    if (this.isLoginMode) {
      this.authService.login(this.user).subscribe(
        () => {
          this.router.navigate(['/tasks']); // Navigate to tasks page on successful login
        },
        (error) => {
          console.error('Login failed:', error);
          this.errorMessage = 'Login failed. Invalid credentials or server error.';
        }
      );
    } else {
      this.authService.signup(this.user).subscribe(
        () => {
          this.isLoginMode = true; // Switch to login mode after successful signup
          this.successMessage = 'Signup successful! Please log in.';
        },
        (error) => {
          console.error('Signup failed:', error);
          this.errorMessage = 'Signup failed. Username might already be taken or server error.';
        }
      );
    }
  }
}
