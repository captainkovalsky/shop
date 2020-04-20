import { Component, OnInit } from '@angular/core';
import {AuthService} from "../../services/auth/auth.service";

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.sass']
})
export class NavigationComponent implements OnInit {
  title = 'frontend';
  public isMenuCollapsed = true;
  constructor(public auth: AuthService) { }

  ngOnInit(): void {
  }

}
