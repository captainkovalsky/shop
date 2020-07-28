import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {map} from "rxjs/operators";

export interface Experience {
  startDate: string;
  endDate: string;
  project: string;
  company: string;
}

export interface CV {
  name: string;
  surname: string;
  experience: Experience[];
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
