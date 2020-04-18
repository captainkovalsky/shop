import { Component } from '@angular/core';
import {AuthService} from "./services/auth/auth.service";
import {BucketService} from "./services/bucket.service";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent {
  title = 'frontend';

  constructor(public auth: AuthService, private bucket: BucketService) {
  }


  ping() {
    this.bucket.ping().subscribe((r) => {
      console.log('response', r)
    })
  }

}
