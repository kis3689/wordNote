import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Observable } from "rxjs";
import { Word } from './word'

@Injectable({
  providedIn: 'root'
})
export class WordService {
  baseUrl: string = '/words'
  readonly headers = new HttpHeaders().set('Content-Type', 'application/json');

  constructor(private http: HttpClient) { }

  getAll(): Observable<Word[]> {
    return this.http.get<Word[]>(this.baseUrl);
  }

  add(wd: Word): Observable<Word> {
    return this.http.post<Word>(this.baseUrl, wd, {headers: this.headers});
  }

  update(wd: Word): Observable<Word> {
    return this.http.put<Word>(
      `${this.baseUrl}/${wd.Id}`, wd, {headers: this.headers}
    );
  }

  delete(id: number): Observable<Word> {
    return this.http.delete<Word>(`${this.baseUrl}/${id}`);
  }
}
