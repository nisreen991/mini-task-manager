import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private apiUrl = 'http://localhost:8080/auth'; // Backend API URL

  constructor(private http: HttpClient) { }

  signup(user: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/signup`, user);
  }

  login(user: any): Observable<string> {
    return this.http.post(`${this.apiUrl}/login`, user, { responseType: 'text' }).pipe(
      tap((token: string) => this.saveToken(token))
    );
  }

  logout(): void {
    localStorage.removeItem('jwt_token');
  }

  public saveToken(token: string): void {
    localStorage.setItem('jwt_token', token);
  }

  public getToken(): string | null {
    return localStorage.getItem('jwt_token');
  }

  public isAuthenticated(): boolean {
    const token = this.getToken();
    // Check if token exists and is not expired (basic check for now)
    return !!token; 
  }
}
