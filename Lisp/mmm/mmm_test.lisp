;;;;;;;;;;;;;;;;;;;;;;;;
;;; pre-testing prep ;;;
;;;;;;;;;;;;;;;;;;;;;;;;

(load "../lisp-unit.lisp")

(use-package :lisp-unit)

(load "mmm.lisp")

(remove-tests :all)

(setq *print-failures* t)

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;; min-mean-max test definitions ;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(define-test test-min-mean-max
    (assert-equal '(5 5 5) (min-mean-max '(5)))
    (assert-equal '(1 3 5) (min-mean-max '(5 4 3 2 1)))
    (assert-equal '(-1 3 8) (min-mean-max '(0 6 8 2 3 -1)))
    (assert-equal '(8 8 8) (min-mean-max '(8 8 8 8 8 8 8 8)))
    (assert-equal '(8 65/8 9) (min-mean-max '(8 8 8 8 8 8 8 9)))
    (assert-equal '(-4 5/2 6) (min-mean-max '(6 6 5 -4 3 2 1 1)))
    (assert-equal '(-3 0 3) (min-mean-max '(0 -1 1 -2 2 -3 3)))
    (assert-equal '(1 2 3) (min-mean-max '(1 2 3)))
    (assert-equal '(-1 -1/2 0) (min-mean-max '(0 -1))))

;;;;;;;;;;;;;;;;;
;;; run tests ;;;
;;;;;;;;;;;;;;;;;

(run-tests :all)
