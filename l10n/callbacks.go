package l10n

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
)

func isLocalizable(scope *gorm.Scope) (isLocalizable bool) {
	_, isLocalizable = reflect.New(scope.GetModelStruct().ModelType).Interface().(Interface)
	return
}

func BeforeQuery(scope *gorm.Scope) {
	if isLocalizable(scope) {
		if str, ok := scope.DB().Get("l10n:locale"); ok {
			if locale, ok := str.(string); ok {
				fmt.Println(locale)
				switch mode, _ := scope.DB().Get("l10n"); mode {
				case "locale":
					// scope.Search.Where("language_code = ?", locale)
				case "global":
					// scope.Search.Where("language_code IS NULL")
				default:
					// scope.Search.Where("language_code = ? OR language_code IS NULL", locale)
				}
			}
		}
	}
}

func BeforeCreate(scope *gorm.Scope) {
	if isLocalizable(scope) {
		// is locale -> set locale
	}
}

func BeforeUpdate(scope *gorm.Scope) {
	if isLocalizable(scope) {
		// is locale -> update localized columns
	}
}

func AfterUpdate(scope *gorm.Scope) {
	if isLocalizable(scope) {
		// is global -> sync colums that need sync
	}
}

func BeforeDelete(scope *gorm.Scope) {
	if isLocalizable(scope) {
		// is locale -> scope.Search.Where("language_code = ?", locale)
	}
}