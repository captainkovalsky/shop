import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { StackIconComponent } from './stack-icon.component';

describe('StackIconComponent', () => {
  let component: StackIconComponent;
  let fixture: ComponentFixture<StackIconComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ StackIconComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(StackIconComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
