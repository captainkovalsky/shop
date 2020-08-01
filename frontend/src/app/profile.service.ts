import {Observable} from "rxjs";
import {Inject, Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {map} from "rxjs/operators";
import {DOCUMENT} from "@angular/common";

import {
  differenceInBusinessDays,
  differenceInCalendarDays,
  differenceInCalendarMonths, differenceInCalendarYears,
  differenceInDays,
  isValid,
  parseISO
} from "date-fns"

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
  companyLogo: string;
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

export interface Period {
  days: number;
  months: number;
}
export interface StackExperience {
  [key: string]: Period;
}


@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
  }


  GetProfile(): Observable<CV> {
    return this.http.get(`${this.document.location.origin}/api/cv`)
      .pipe(
        map(r => r as CV)
      );
  }

  GetStackExperience(experience: Experience[]): StackExperience | null {
    experience.forEach(exp => {
      if (!isValid(parseISO(exp.startDate))) {
        throw new Error(`${exp.startDate} is not valid`)
      }
      if (!isValid(parseISO(exp.endDate))) {
        throw new Error(`${exp.endDate} is not valid`)
      }
    });


    const response: StackExperience = experience.reduce((prev, curr) => {
      const endDate = parseISO(curr.endDate),
        startDate = parseISO(curr.startDate);


      const days = differenceInCalendarDays(endDate, startDate)
      const month = differenceInCalendarMonths(endDate, startDate)

      curr.stack.forEach((tech) => {
        console.log(`${curr.company} - ${tech}`)
        // @ts-ignore
        if (prev[tech] != undefined) {
          console.log(`has tech in ${prev}`);
          // @ts-ignore
          (prev[tech] as Period).days += days;
          // @ts-ignore
          (prev[tech] as Period).months += month;
        } else {
          console.log(`NO tech in ${prev}`);
          // @ts-ignore
          prev[tech] = {};
          // @ts-ignore
          (prev[tech] as Period).days = days;
          // @ts-ignore
          (prev[tech] as Period).months = month;
        }
      })

      return prev;
    }, {});


    return response;
  }
}
