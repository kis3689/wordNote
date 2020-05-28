import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WordnoteComponent } from './wordnote.component';

describe('WordnoteComponent', () => {
  let component: WordnoteComponent;
  let fixture: ComponentFixture<WordnoteComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WordnoteComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WordnoteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
