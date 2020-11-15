import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { StackPageComponent } from './stack-page.component';

describe('StackPageComponent', () => {
  let component: StackPageComponent;
  let fixture: ComponentFixture<StackPageComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ StackPageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(StackPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
