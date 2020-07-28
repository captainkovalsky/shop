import {Component, OnInit} from '@angular/core';
import {CV, ProfileService} from "./profile.service";

@Component({
  selector: 'app-root',
  template: `
<section class="section">
<pre hidden>{{cv | json}}</pre>
    <router-outlet></router-outlet>
</section>
  `,
  styles: []
})
export class AppComponent implements OnInit {
  title = 'CV';
  cv: CV | null = null;

  constructor(private profileSrv: ProfileService) {

  }


  ngOnInit(): void {
    this.profileSrv.GetProfile().subscribe((profile: CV) => {
      this.cv = profile;
    })
  }


}
