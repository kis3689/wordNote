import { TestBed } from '@angular/core/testing';

import { WordnoteService } from './wordnote.service';

describe('WordnoteService', () => {
  let service: WordnoteService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(WordnoteService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
