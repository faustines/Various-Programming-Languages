;;;;;;;;;;;;;;;;;;;;;;;;
;;; pre-testing prep ;;;
;;;;;;;;;;;;;;;;;;;;;;;;

(load "../lisp-unit.lisp")

(use-package :lisp-unit)

(load "match.lisp")

(remove-tests :all)

(setq *print-failures* t)

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;; match test definitions ;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(define-test test-match
    (assert-equal NIL (match '(THINK ABOUT THAT) '(THINK THAT ABOUT)))
    (assert-equal T (match '(THINK ABOUT THAT) '(THINK ABOUT THAT)))
    (assert-equal NIL (match '(THINK ABOUT THAT) NIL))
    (assert-equal T (match NIL 'NIL))
    (assert-equal NIL (match NIL '(NIL)))
    (assert-equal NIL (match '(NIL NIL) '(NIL)))
    (assert-equal T (match '(NIL NIL) '(NIL NIL)))
    (assert-equal NIL (match '(A A) '(A)))
    (assert-equal T (match '(A ?) '(A A)))
    (assert-equal T (match '(A NIL ?) '(A NIL NIL)))
    (assert-equal T (match '(? ? ? ?) '(WHAT TIME IS IT)))
    (assert-equal NIL (match '(? ? ? ? ?) '(WHAT TIME IS IT NOW AH)))
    (assert-equal NIL (match '(? ? ? ? ?) '(WHAT TIME IS IT)))
    (assert-equal NIL (match '(? COLOR APPLE) '(COLOR COLOR APPLE NIL)))
    (assert-equal T (match '(THIS TABLE IS ROUND ?) '(THIS TABLE IS ROUND OR)))
    (assert-equal T (match '(THIS !) '(THIS TABLE IS ROUND)))
    (assert-equal T (match '(THIS ! IS ROUND) '(THIS TABLE IS ROUND)))
    (assert-equal NIL (match '(THIS TABLE !) '(THIS THIS TABLE IS ROUND)))
    (assert-equal NIL (match '(! TABLE IS ROUND) '(THIS TABLE IS ROUND OR)))
    (assert-equal NIL (match '(! TABLE IS ROUND) '(THIS TABLE)))
    (assert-equal T (match '(THIS ! ROUND) '(THIS TABLE IS ROUND)))
    (assert-equal NIL (match '(IS THAT FINE ?) '(THAT FINE TOO)))
    (assert-equal T (match '(!) '(NO MATTER WHAT)))
    (assert-equal T (match '(! ! !) '(NO MATTER WHAT)))
    (assert-equal NIL (match '(! ! ! !) '(NO MATTER WHAT)))
    (assert-equal T (match '(! BIG ! ?) '(A YELLOW BIG WOODEN BAG)))
    (assert-equal NIL (match '(! BIG ! ?) '(A YELLOW BIG BAG)))
    (assert-equal T (match '(! ! ? ! !) '(HOW ARE YOU TODAY AH)))
    (assert-equal NIL (match '(! ! ! ! !) '(HOW ARE YOU TODAY)))
    (assert-equal T (match '(! ? ? FUNNY !) '(IS IT REALLY FUNNY OR WHAT)))
    (assert-equal NIL (match '(! ? ? FUNNY !) '(IS IT REALLY FUNNY)))
    (assert-equal NIL (match '(? ? ! ? A) '(A A A A)))
    (assert-equal T (match '(? ? ! ? A !) '(A A A A A NIL)))
    (assert-equal T (match '(! B ? B A B A ! ? !) '(B B A B A B A B A B A)))
    (assert-equal NIL (match '(! B ? A B A B ! ? !) '(B B A B A B A B A B))))

;;;;;;;;;;;;;;;;;
;;; run tests ;;;
;;;;;;;;;;;;;;;;;

(run-tests :all)
