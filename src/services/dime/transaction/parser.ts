export interface Parser<T> {
  toJson(): T;
  save(): void;
}
