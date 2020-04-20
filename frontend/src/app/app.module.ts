import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FbAuthComponent } from './components/fb-auth/fb-auth.component';
import { CallbackComponent } from './components/callback/callback.component';
import {HTTP_INTERCEPTORS, HttpClientModule} from "@angular/common/http";
import {TokenInterceptor} from "./services/interceptors/token.interceptor";
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import { NavigationComponent } from './components/navigation/navigation.component';
import { LoginComponent } from './components/login/login.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { DeployComponent } from './components/deploy/deploy.component';
import { DataConfigurationComponent } from './components/data-configuration/data-configuration.component';
@NgModule({
  declarations: [
    AppComponent,
    FbAuthComponent,
    CallbackComponent,
    NavigationComponent,
    LoginComponent,
    DashboardComponent,
    DeployComponent,
    DataConfigurationComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    NgbModule
  ],
    providers: [{
      provide: HTTP_INTERCEPTORS,
      useClass: TokenInterceptor,
      multi: true
    }],
  bootstrap: [AppComponent]
})
export class AppModule { }
