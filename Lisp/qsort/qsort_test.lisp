;;;;;;;;;;;;;;;;;;;;;;;;
;;; pre-testing prep ;;;
;;;;;;;;;;;;;;;;;;;;;;;;

(load "../lisp-unit.lisp")

(use-package :lisp-unit)

(load "qsort.lisp")

(remove-tests :all)

(setq *print-failures* t)


;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;; pivot test definitions ;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(define-test test-pivot
    (assert-equal '(NIL NIL) (pivot 2 'NIL))
    (assert-equal '((4 3 1) (6 7 9)) (pivot 5 '(6 4 3 1 7 9)))
    (assert-equal '((2 1) (3 5 4)) (pivot 3 '(3 2 5 1 4)))
    (assert-equal '((1 2 3 5 4 5) NIL) (pivot 6 '(1 2 3 5 4 5)))
    (assert-equal '(NIL (8 9 12 11 15)) (pivot 8 '(8 9 12 11 15)))
    (assert-equal '((4 5 2 9) (21 11 14 15 13)) (pivot 10 '(21 11 4 14 5 15 13 2 9)))
    (assert-equal '((-6 4 -3 1 -9) (7)) (pivot 5 '(-6 4 -3 1 7 -9)))
    (assert-equal '((-5 -4) (3 2 1)) (pivot -3 '(3 2 -5 1 -4)))
    (assert-equal '((-1 2 3 5 4 -5) NIL) (pivot 6 '(-1 2 3 5 4 -5)))
    (assert-equal '((-9 -12 -11 -15) (-8)) (pivot -8 '(-8 -9 -12 -11 -15))))
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;; quicksort test definitions ;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(define-test test-quicksort
    (assert-equal 'NIL (quicksort 'NIL))
    (assert-equal '(1 3 4 6 7 9) (quicksort '(6 4 3 1 7 9)))
    (assert-equal '(-2 1 3 4 5 8) (quicksort '(8 1 -2 3 5 4)))
    (assert-equal '(8 9 11 12 15) (quicksort '(8 9 12 11 15)))
    (assert-equal '(-14 2 4 5 9 11 13 15 21) (quicksort '(21 11 4 -14 5 15 13 2 9)))
    (assert-equal '(-14 -13 -11 -5 2 4 9 15 21) (quicksort '(21 -11 4 -14 -5 15 -13 2 9)))
    (assert-equal '(0 2 4 5 9 11 13 14 15 21) (quicksort '(21 11 4 0 14 5 15 13 2 9)))
    (assert-equal '(8) (quicksort '(8)))
    (assert-equal '(2 8) (quicksort '(8 2)))
    (assert-equal '(-8 2 8) (quicksort '(-8 2 8))))

;;;;;;;;;;;;;;;;;
;;; run tests ;;;
;;;;;;;;;;;;;;;;;

(run-tests :all)
