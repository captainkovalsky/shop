import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {map} from "rxjs/operators";

export interface Contacts {
  skype: string;
  phone: string;
  email: string;
  linkedin: string;
  github: string;
}

export interface Experience {
  startDate: string;
  endDate: string;
  role: string;
  project: string;
  company: string;
  projectSite: string;
  frontendSize: number;
  backendSize: number;
  stack: string[];
  responsibilities: any[];
  qaSize?: number;
}

export interface Question {
  lang: string;
  text: string;
}

export interface Answer {
  lang: string;
  text: string;
}

export interface Faq {
  question: Question[];
  answer: Answer[];
  company: string;
  project: string;
}

export interface Hourly {
  amount: number;
  currency: string;
}

export interface Monthly {
  amount: number;
  currency: string;
}

export interface Compensation {
  hourly: Hourly;
  monthly: Monthly;
}

export interface CV {
  name: string;
  surname: string;
  contacts: Contacts;
  summary: string[];
  experience: Experience[];
  faq: Faq[];
  compensation: Compensation;
}

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor(
    private http: HttpClient
  ) {
  }


  GetProfile(): Observable<CV> {
    return this.http.get('http://localhost:9001/api/cv')
      .pipe(
        map(r => r as CV)
      );
  }
}
