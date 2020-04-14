import {Component, OnInit} from '@angular/core';
import {environment} from "../../environments/environment";

declare var FB: any;

@Component({
  selector: 'app-fb-auth',
  template: `
<container-element [ngSwitch]="auth?.status">
  <p *ngSwitchCase="'connected'">
  <button type="button" (click)="logout();">
  Logout
</button>
</p>
  <p *ngSwitchDefault>
    <button
  type="button"
  (click)="submitLogin();">Login with facebook</button>
</p>
</container-element>
  `,
  styleUrls: ['./fb-auth.component.sass']
})
export class FbAuthComponent implements OnInit {
  public auth = null;

  constructor() {
  }

  ngOnInit(): void {
    (window as any).fbAsyncInit = function () {
      FB.init({
        appId: environment.facebookAppId,
        cookie: true,
        xfbml: true,
        version: 'v3.1'
      });
      FB.AppEvents.logPageView();
    };

    (function (d, s, id) {
      var js, fjs = d.getElementsByTagName(s)[0];
      if (d.getElementById(id)) {
        return;
      }
      js = d.createElement(s);
      js.id = id;
      js.src = "https://connect.facebook.net/en_US/sdk.js";
      fjs.parentNode.insertBefore(js, fjs);
    }(document, 'script', 'facebook-jssdk'));
  }

  logout() {
    FB.logout((response) => {
      this.auth = null;
      console.log('logout', response)
    });
  }

  submitLogin() {
    console.log("submit login to facebook");
    FB.login((response) => {
      this.auth = response;
      if (response.authResponse) {
        //login success
        //login success code here
        //redirect to home page
      } else {
        console.log('User login failed');
      }
    });

  }

}
