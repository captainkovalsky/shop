import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {environment} from "../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class BucketService {

  constructor(private httpClient: HttpClient) { }


  ping() {
    return this.httpClient.get(environment.gateway + '/ping');
  }

}
