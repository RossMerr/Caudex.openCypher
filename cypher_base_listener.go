// Generated from Cypher.g4 by ANTLR 4.7.

package openCypher // Cypher
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCypherListener is a complete listener for a parse tree produced by CypherParser.
type BaseCypherListener struct{}

var _ CypherListener = &BaseCypherListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCypherListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCypherListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCypherListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCypherListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCypher is called when production cypher is entered.
func (s *BaseCypherListener) EnterCypher(ctx *CypherContext) {}

// ExitCypher is called when production cypher is exited.
func (s *BaseCypherListener) ExitCypher(ctx *CypherContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCypherListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCypherListener) ExitStatement(ctx *StatementContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseCypherListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseCypherListener) ExitQuery(ctx *QueryContext) {}

// EnterRegularQuery is called when production regularQuery is entered.
func (s *BaseCypherListener) EnterRegularQuery(ctx *RegularQueryContext) {}

// ExitRegularQuery is called when production regularQuery is exited.
func (s *BaseCypherListener) ExitRegularQuery(ctx *RegularQueryContext) {}

// EnterUnion is called when production union is entered.
func (s *BaseCypherListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production union is exited.
func (s *BaseCypherListener) ExitUnion(ctx *UnionContext) {}

// EnterSingleQuery is called when production singleQuery is entered.
func (s *BaseCypherListener) EnterSingleQuery(ctx *SingleQueryContext) {}

// ExitSingleQuery is called when production singleQuery is exited.
func (s *BaseCypherListener) ExitSingleQuery(ctx *SingleQueryContext) {}

// EnterSinglePartQuery is called when production singlePartQuery is entered.
func (s *BaseCypherListener) EnterSinglePartQuery(ctx *SinglePartQueryContext) {}

// ExitSinglePartQuery is called when production singlePartQuery is exited.
func (s *BaseCypherListener) ExitSinglePartQuery(ctx *SinglePartQueryContext) {}

// EnterReadOnlyEnd is called when production readOnlyEnd is entered.
func (s *BaseCypherListener) EnterReadOnlyEnd(ctx *ReadOnlyEndContext) {}

// ExitReadOnlyEnd is called when production readOnlyEnd is exited.
func (s *BaseCypherListener) ExitReadOnlyEnd(ctx *ReadOnlyEndContext) {}

// EnterReadUpdateEnd is called when production readUpdateEnd is entered.
func (s *BaseCypherListener) EnterReadUpdateEnd(ctx *ReadUpdateEndContext) {}

// ExitReadUpdateEnd is called when production readUpdateEnd is exited.
func (s *BaseCypherListener) ExitReadUpdateEnd(ctx *ReadUpdateEndContext) {}

// EnterUpdatingEnd is called when production updatingEnd is entered.
func (s *BaseCypherListener) EnterUpdatingEnd(ctx *UpdatingEndContext) {}

// ExitUpdatingEnd is called when production updatingEnd is exited.
func (s *BaseCypherListener) ExitUpdatingEnd(ctx *UpdatingEndContext) {}

// EnterMultiPartQuery is called when production multiPartQuery is entered.
func (s *BaseCypherListener) EnterMultiPartQuery(ctx *MultiPartQueryContext) {}

// ExitMultiPartQuery is called when production multiPartQuery is exited.
func (s *BaseCypherListener) ExitMultiPartQuery(ctx *MultiPartQueryContext) {}

// EnterReadPart is called when production readPart is entered.
func (s *BaseCypherListener) EnterReadPart(ctx *ReadPartContext) {}

// ExitReadPart is called when production readPart is exited.
func (s *BaseCypherListener) ExitReadPart(ctx *ReadPartContext) {}

// EnterUpdatingPart is called when production updatingPart is entered.
func (s *BaseCypherListener) EnterUpdatingPart(ctx *UpdatingPartContext) {}

// ExitUpdatingPart is called when production updatingPart is exited.
func (s *BaseCypherListener) ExitUpdatingPart(ctx *UpdatingPartContext) {}

// EnterUpdatingStartClause is called when production updatingStartClause is entered.
func (s *BaseCypherListener) EnterUpdatingStartClause(ctx *UpdatingStartClauseContext) {}

// ExitUpdatingStartClause is called when production updatingStartClause is exited.
func (s *BaseCypherListener) ExitUpdatingStartClause(ctx *UpdatingStartClauseContext) {}

// EnterUpdatingClause is called when production updatingClause is entered.
func (s *BaseCypherListener) EnterUpdatingClause(ctx *UpdatingClauseContext) {}

// ExitUpdatingClause is called when production updatingClause is exited.
func (s *BaseCypherListener) ExitUpdatingClause(ctx *UpdatingClauseContext) {}

// EnterReadingClause is called when production readingClause is entered.
func (s *BaseCypherListener) EnterReadingClause(ctx *ReadingClauseContext) {}

// ExitReadingClause is called when production readingClause is exited.
func (s *BaseCypherListener) ExitReadingClause(ctx *ReadingClauseContext) {}

// EnterMatch is called when production match is entered.
func (s *BaseCypherListener) EnterMatch(ctx *MatchContext) {}

// ExitMatch is called when production match is exited.
func (s *BaseCypherListener) ExitMatch(ctx *MatchContext) {}

// EnterUnwind is called when production unwind is entered.
func (s *BaseCypherListener) EnterUnwind(ctx *UnwindContext) {}

// ExitUnwind is called when production unwind is exited.
func (s *BaseCypherListener) ExitUnwind(ctx *UnwindContext) {}

// EnterMerge is called when production merge is entered.
func (s *BaseCypherListener) EnterMerge(ctx *MergeContext) {}

// ExitMerge is called when production merge is exited.
func (s *BaseCypherListener) ExitMerge(ctx *MergeContext) {}

// EnterMergeAction is called when production mergeAction is entered.
func (s *BaseCypherListener) EnterMergeAction(ctx *MergeActionContext) {}

// ExitMergeAction is called when production mergeAction is exited.
func (s *BaseCypherListener) ExitMergeAction(ctx *MergeActionContext) {}

// EnterCreate is called when production create is entered.
func (s *BaseCypherListener) EnterCreate(ctx *CreateContext) {}

// ExitCreate is called when production create is exited.
func (s *BaseCypherListener) ExitCreate(ctx *CreateContext) {}

// EnterSet is called when production set is entered.
func (s *BaseCypherListener) EnterSet(ctx *SetContext) {}

// ExitSet is called when production set is exited.
func (s *BaseCypherListener) ExitSet(ctx *SetContext) {}

// EnterSetItem is called when production setItem is entered.
func (s *BaseCypherListener) EnterSetItem(ctx *SetItemContext) {}

// ExitSetItem is called when production setItem is exited.
func (s *BaseCypherListener) ExitSetItem(ctx *SetItemContext) {}

// EnterDelete is called when production delete is entered.
func (s *BaseCypherListener) EnterDelete(ctx *DeleteContext) {}

// ExitDelete is called when production delete is exited.
func (s *BaseCypherListener) ExitDelete(ctx *DeleteContext) {}

// EnterRemove is called when production remove is entered.
func (s *BaseCypherListener) EnterRemove(ctx *RemoveContext) {}

// ExitRemove is called when production remove is exited.
func (s *BaseCypherListener) ExitRemove(ctx *RemoveContext) {}

// EnterRemoveItem is called when production removeItem is entered.
func (s *BaseCypherListener) EnterRemoveItem(ctx *RemoveItemContext) {}

// ExitRemoveItem is called when production removeItem is exited.
func (s *BaseCypherListener) ExitRemoveItem(ctx *RemoveItemContext) {}

// EnterInQueryCall is called when production inQueryCall is entered.
func (s *BaseCypherListener) EnterInQueryCall(ctx *InQueryCallContext) {}

// ExitInQueryCall is called when production inQueryCall is exited.
func (s *BaseCypherListener) ExitInQueryCall(ctx *InQueryCallContext) {}

// EnterStandaloneCall is called when production standaloneCall is entered.
func (s *BaseCypherListener) EnterStandaloneCall(ctx *StandaloneCallContext) {}

// ExitStandaloneCall is called when production standaloneCall is exited.
func (s *BaseCypherListener) ExitStandaloneCall(ctx *StandaloneCallContext) {}

// EnterYieldItems is called when production yieldItems is entered.
func (s *BaseCypherListener) EnterYieldItems(ctx *YieldItemsContext) {}

// ExitYieldItems is called when production yieldItems is exited.
func (s *BaseCypherListener) ExitYieldItems(ctx *YieldItemsContext) {}

// EnterYieldItem is called when production yieldItem is entered.
func (s *BaseCypherListener) EnterYieldItem(ctx *YieldItemContext) {}

// ExitYieldItem is called when production yieldItem is exited.
func (s *BaseCypherListener) ExitYieldItem(ctx *YieldItemContext) {}

// EnterWith is called when production with is entered.
func (s *BaseCypherListener) EnterWith(ctx *WithContext) {}

// ExitWith is called when production with is exited.
func (s *BaseCypherListener) ExitWith(ctx *WithContext) {}

// EnterReturn is called when production return is entered.
func (s *BaseCypherListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production return is exited.
func (s *BaseCypherListener) ExitReturn(ctx *ReturnContext) {}

// EnterReturnBody is called when production returnBody is entered.
func (s *BaseCypherListener) EnterReturnBody(ctx *ReturnBodyContext) {}

// ExitReturnBody is called when production returnBody is exited.
func (s *BaseCypherListener) ExitReturnBody(ctx *ReturnBodyContext) {}

// EnterReturnItems is called when production returnItems is entered.
func (s *BaseCypherListener) EnterReturnItems(ctx *ReturnItemsContext) {}

// ExitReturnItems is called when production returnItems is exited.
func (s *BaseCypherListener) ExitReturnItems(ctx *ReturnItemsContext) {}

// EnterReturnItem is called when production returnItem is entered.
func (s *BaseCypherListener) EnterReturnItem(ctx *ReturnItemContext) {}

// ExitReturnItem is called when production returnItem is exited.
func (s *BaseCypherListener) ExitReturnItem(ctx *ReturnItemContext) {}

// EnterOrder is called when production order is entered.
func (s *BaseCypherListener) EnterOrder(ctx *OrderContext) {}

// ExitOrder is called when production order is exited.
func (s *BaseCypherListener) ExitOrder(ctx *OrderContext) {}

// EnterSkip is called when production skip is entered.
func (s *BaseCypherListener) EnterSkip(ctx *SkipContext) {}

// ExitSkip is called when production skip is exited.
func (s *BaseCypherListener) ExitSkip(ctx *SkipContext) {}

// EnterLimit is called when production limit is entered.
func (s *BaseCypherListener) EnterLimit(ctx *LimitContext) {}

// ExitLimit is called when production limit is exited.
func (s *BaseCypherListener) ExitLimit(ctx *LimitContext) {}

// EnterSortItem is called when production sortItem is entered.
func (s *BaseCypherListener) EnterSortItem(ctx *SortItemContext) {}

// ExitSortItem is called when production sortItem is exited.
func (s *BaseCypherListener) ExitSortItem(ctx *SortItemContext) {}

// EnterWhere is called when production where is entered.
func (s *BaseCypherListener) EnterWhere(ctx *WhereContext) {}

// ExitWhere is called when production where is exited.
func (s *BaseCypherListener) ExitWhere(ctx *WhereContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseCypherListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseCypherListener) ExitPattern(ctx *PatternContext) {}

// EnterPatternPart is called when production patternPart is entered.
func (s *BaseCypherListener) EnterPatternPart(ctx *PatternPartContext) {}

// ExitPatternPart is called when production patternPart is exited.
func (s *BaseCypherListener) ExitPatternPart(ctx *PatternPartContext) {}

// EnterAnonymousPatternPart is called when production anonymousPatternPart is entered.
func (s *BaseCypherListener) EnterAnonymousPatternPart(ctx *AnonymousPatternPartContext) {}

// ExitAnonymousPatternPart is called when production anonymousPatternPart is exited.
func (s *BaseCypherListener) ExitAnonymousPatternPart(ctx *AnonymousPatternPartContext) {}

// EnterPatternElement is called when production patternElement is entered.
func (s *BaseCypherListener) EnterPatternElement(ctx *PatternElementContext) {}

// ExitPatternElement is called when production patternElement is exited.
func (s *BaseCypherListener) ExitPatternElement(ctx *PatternElementContext) {}

// EnterNodePattern is called when production nodePattern is entered.
func (s *BaseCypherListener) EnterNodePattern(ctx *NodePatternContext) {}

// ExitNodePattern is called when production nodePattern is exited.
func (s *BaseCypherListener) ExitNodePattern(ctx *NodePatternContext) {}

// EnterPatternElementChain is called when production patternElementChain is entered.
func (s *BaseCypherListener) EnterPatternElementChain(ctx *PatternElementChainContext) {}

// ExitPatternElementChain is called when production patternElementChain is exited.
func (s *BaseCypherListener) ExitPatternElementChain(ctx *PatternElementChainContext) {}

// EnterRelationshipPattern is called when production relationshipPattern is entered.
func (s *BaseCypherListener) EnterRelationshipPattern(ctx *RelationshipPatternContext) {}

// ExitRelationshipPattern is called when production relationshipPattern is exited.
func (s *BaseCypherListener) ExitRelationshipPattern(ctx *RelationshipPatternContext) {}

// EnterRelationshipDetail is called when production relationshipDetail is entered.
func (s *BaseCypherListener) EnterRelationshipDetail(ctx *RelationshipDetailContext) {}

// ExitRelationshipDetail is called when production relationshipDetail is exited.
func (s *BaseCypherListener) ExitRelationshipDetail(ctx *RelationshipDetailContext) {}

// EnterProperties is called when production properties is entered.
func (s *BaseCypherListener) EnterProperties(ctx *PropertiesContext) {}

// ExitProperties is called when production properties is exited.
func (s *BaseCypherListener) ExitProperties(ctx *PropertiesContext) {}

// EnterRelationshipTypes is called when production relationshipTypes is entered.
func (s *BaseCypherListener) EnterRelationshipTypes(ctx *RelationshipTypesContext) {}

// ExitRelationshipTypes is called when production relationshipTypes is exited.
func (s *BaseCypherListener) ExitRelationshipTypes(ctx *RelationshipTypesContext) {}

// EnterNodeLabels is called when production nodeLabels is entered.
func (s *BaseCypherListener) EnterNodeLabels(ctx *NodeLabelsContext) {}

// ExitNodeLabels is called when production nodeLabels is exited.
func (s *BaseCypherListener) ExitNodeLabels(ctx *NodeLabelsContext) {}

// EnterNodeLabel is called when production nodeLabel is entered.
func (s *BaseCypherListener) EnterNodeLabel(ctx *NodeLabelContext) {}

// ExitNodeLabel is called when production nodeLabel is exited.
func (s *BaseCypherListener) ExitNodeLabel(ctx *NodeLabelContext) {}

// EnterRangeLiteral is called when production rangeLiteral is entered.
func (s *BaseCypherListener) EnterRangeLiteral(ctx *RangeLiteralContext) {}

// ExitRangeLiteral is called when production rangeLiteral is exited.
func (s *BaseCypherListener) ExitRangeLiteral(ctx *RangeLiteralContext) {}

// EnterLabelName is called when production labelName is entered.
func (s *BaseCypherListener) EnterLabelName(ctx *LabelNameContext) {}

// ExitLabelName is called when production labelName is exited.
func (s *BaseCypherListener) ExitLabelName(ctx *LabelNameContext) {}

// EnterRelTypeName is called when production relTypeName is entered.
func (s *BaseCypherListener) EnterRelTypeName(ctx *RelTypeNameContext) {}

// ExitRelTypeName is called when production relTypeName is exited.
func (s *BaseCypherListener) ExitRelTypeName(ctx *RelTypeNameContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCypherListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCypherListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseCypherListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseCypherListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterXorExpression is called when production xorExpression is entered.
func (s *BaseCypherListener) EnterXorExpression(ctx *XorExpressionContext) {}

// ExitXorExpression is called when production xorExpression is exited.
func (s *BaseCypherListener) ExitXorExpression(ctx *XorExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseCypherListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseCypherListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseCypherListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseCypherListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterComparisonExpression is called when production comparisonExpression is entered.
func (s *BaseCypherListener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production comparisonExpression is exited.
func (s *BaseCypherListener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterAddOrSubtractExpression is called when production addOrSubtractExpression is entered.
func (s *BaseCypherListener) EnterAddOrSubtractExpression(ctx *AddOrSubtractExpressionContext) {}

// ExitAddOrSubtractExpression is called when production addOrSubtractExpression is exited.
func (s *BaseCypherListener) ExitAddOrSubtractExpression(ctx *AddOrSubtractExpressionContext) {}

// EnterMultiplyDivideModuloExpression is called when production multiplyDivideModuloExpression is entered.
func (s *BaseCypherListener) EnterMultiplyDivideModuloExpression(ctx *MultiplyDivideModuloExpressionContext) {
}

// ExitMultiplyDivideModuloExpression is called when production multiplyDivideModuloExpression is exited.
func (s *BaseCypherListener) ExitMultiplyDivideModuloExpression(ctx *MultiplyDivideModuloExpressionContext) {
}

// EnterPowerOfExpression is called when production powerOfExpression is entered.
func (s *BaseCypherListener) EnterPowerOfExpression(ctx *PowerOfExpressionContext) {}

// ExitPowerOfExpression is called when production powerOfExpression is exited.
func (s *BaseCypherListener) ExitPowerOfExpression(ctx *PowerOfExpressionContext) {}

// EnterUnaryAddOrSubtractExpression is called when production unaryAddOrSubtractExpression is entered.
func (s *BaseCypherListener) EnterUnaryAddOrSubtractExpression(ctx *UnaryAddOrSubtractExpressionContext) {
}

// ExitUnaryAddOrSubtractExpression is called when production unaryAddOrSubtractExpression is exited.
func (s *BaseCypherListener) ExitUnaryAddOrSubtractExpression(ctx *UnaryAddOrSubtractExpressionContext) {
}

// EnterStringListNullOperatorExpression is called when production stringListNullOperatorExpression is entered.
func (s *BaseCypherListener) EnterStringListNullOperatorExpression(ctx *StringListNullOperatorExpressionContext) {
}

// ExitStringListNullOperatorExpression is called when production stringListNullOperatorExpression is exited.
func (s *BaseCypherListener) ExitStringListNullOperatorExpression(ctx *StringListNullOperatorExpressionContext) {
}

// EnterPropertyOrLabelsExpression is called when production propertyOrLabelsExpression is entered.
func (s *BaseCypherListener) EnterPropertyOrLabelsExpression(ctx *PropertyOrLabelsExpressionContext) {}

// ExitPropertyOrLabelsExpression is called when production propertyOrLabelsExpression is exited.
func (s *BaseCypherListener) ExitPropertyOrLabelsExpression(ctx *PropertyOrLabelsExpressionContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseCypherListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseCypherListener) ExitAtom(ctx *AtomContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseCypherListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseCypherListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseCypherListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseCypherListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterListLiteral is called when production listLiteral is entered.
func (s *BaseCypherListener) EnterListLiteral(ctx *ListLiteralContext) {}

// ExitListLiteral is called when production listLiteral is exited.
func (s *BaseCypherListener) ExitListLiteral(ctx *ListLiteralContext) {}

// EnterPartialComparisonExpression is called when production partialComparisonExpression is entered.
func (s *BaseCypherListener) EnterPartialComparisonExpression(ctx *PartialComparisonExpressionContext) {
}

// ExitPartialComparisonExpression is called when production partialComparisonExpression is exited.
func (s *BaseCypherListener) ExitPartialComparisonExpression(ctx *PartialComparisonExpressionContext) {
}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseCypherListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseCypherListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterRelationshipsPattern is called when production relationshipsPattern is entered.
func (s *BaseCypherListener) EnterRelationshipsPattern(ctx *RelationshipsPatternContext) {}

// ExitRelationshipsPattern is called when production relationshipsPattern is exited.
func (s *BaseCypherListener) ExitRelationshipsPattern(ctx *RelationshipsPatternContext) {}

// EnterFilterExpression is called when production filterExpression is entered.
func (s *BaseCypherListener) EnterFilterExpression(ctx *FilterExpressionContext) {}

// ExitFilterExpression is called when production filterExpression is exited.
func (s *BaseCypherListener) ExitFilterExpression(ctx *FilterExpressionContext) {}

// EnterIdInColl is called when production idInColl is entered.
func (s *BaseCypherListener) EnterIdInColl(ctx *IdInCollContext) {}

// ExitIdInColl is called when production idInColl is exited.
func (s *BaseCypherListener) ExitIdInColl(ctx *IdInCollContext) {}

// EnterFunctionInvocation is called when production functionInvocation is entered.
func (s *BaseCypherListener) EnterFunctionInvocation(ctx *FunctionInvocationContext) {}

// ExitFunctionInvocation is called when production functionInvocation is exited.
func (s *BaseCypherListener) ExitFunctionInvocation(ctx *FunctionInvocationContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseCypherListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseCypherListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterExplicitProcedureInvocation is called when production explicitProcedureInvocation is entered.
func (s *BaseCypherListener) EnterExplicitProcedureInvocation(ctx *ExplicitProcedureInvocationContext) {
}

// ExitExplicitProcedureInvocation is called when production explicitProcedureInvocation is exited.
func (s *BaseCypherListener) ExitExplicitProcedureInvocation(ctx *ExplicitProcedureInvocationContext) {
}

// EnterImplicitProcedureInvocation is called when production implicitProcedureInvocation is entered.
func (s *BaseCypherListener) EnterImplicitProcedureInvocation(ctx *ImplicitProcedureInvocationContext) {
}

// ExitImplicitProcedureInvocation is called when production implicitProcedureInvocation is exited.
func (s *BaseCypherListener) ExitImplicitProcedureInvocation(ctx *ImplicitProcedureInvocationContext) {
}

// EnterProcedureResultField is called when production procedureResultField is entered.
func (s *BaseCypherListener) EnterProcedureResultField(ctx *ProcedureResultFieldContext) {}

// ExitProcedureResultField is called when production procedureResultField is exited.
func (s *BaseCypherListener) ExitProcedureResultField(ctx *ProcedureResultFieldContext) {}

// EnterProcedureName is called when production procedureName is entered.
func (s *BaseCypherListener) EnterProcedureName(ctx *ProcedureNameContext) {}

// ExitProcedureName is called when production procedureName is exited.
func (s *BaseCypherListener) ExitProcedureName(ctx *ProcedureNameContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseCypherListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseCypherListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterListComprehension is called when production listComprehension is entered.
func (s *BaseCypherListener) EnterListComprehension(ctx *ListComprehensionContext) {}

// ExitListComprehension is called when production listComprehension is exited.
func (s *BaseCypherListener) ExitListComprehension(ctx *ListComprehensionContext) {}

// EnterPatternComprehension is called when production patternComprehension is entered.
func (s *BaseCypherListener) EnterPatternComprehension(ctx *PatternComprehensionContext) {}

// ExitPatternComprehension is called when production patternComprehension is exited.
func (s *BaseCypherListener) ExitPatternComprehension(ctx *PatternComprehensionContext) {}

// EnterPropertyLookup is called when production propertyLookup is entered.
func (s *BaseCypherListener) EnterPropertyLookup(ctx *PropertyLookupContext) {}

// ExitPropertyLookup is called when production propertyLookup is exited.
func (s *BaseCypherListener) ExitPropertyLookup(ctx *PropertyLookupContext) {}

// EnterCaseExpression is called when production caseExpression is entered.
func (s *BaseCypherListener) EnterCaseExpression(ctx *CaseExpressionContext) {}

// ExitCaseExpression is called when production caseExpression is exited.
func (s *BaseCypherListener) ExitCaseExpression(ctx *CaseExpressionContext) {}

// EnterCaseAlternatives is called when production caseAlternatives is entered.
func (s *BaseCypherListener) EnterCaseAlternatives(ctx *CaseAlternativesContext) {}

// ExitCaseAlternatives is called when production caseAlternatives is exited.
func (s *BaseCypherListener) ExitCaseAlternatives(ctx *CaseAlternativesContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseCypherListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseCypherListener) ExitVariable(ctx *VariableContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseCypherListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseCypherListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterMapLiteral is called when production mapLiteral is entered.
func (s *BaseCypherListener) EnterMapLiteral(ctx *MapLiteralContext) {}

// ExitMapLiteral is called when production mapLiteral is exited.
func (s *BaseCypherListener) ExitMapLiteral(ctx *MapLiteralContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseCypherListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseCypherListener) ExitParameter(ctx *ParameterContext) {}

// EnterPropertyExpression is called when production propertyExpression is entered.
func (s *BaseCypherListener) EnterPropertyExpression(ctx *PropertyExpressionContext) {}

// ExitPropertyExpression is called when production propertyExpression is exited.
func (s *BaseCypherListener) ExitPropertyExpression(ctx *PropertyExpressionContext) {}

// EnterPropertyKeyName is called when production propertyKeyName is entered.
func (s *BaseCypherListener) EnterPropertyKeyName(ctx *PropertyKeyNameContext) {}

// ExitPropertyKeyName is called when production propertyKeyName is exited.
func (s *BaseCypherListener) ExitPropertyKeyName(ctx *PropertyKeyNameContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseCypherListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseCypherListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterDoubleLiteral is called when production doubleLiteral is entered.
func (s *BaseCypherListener) EnterDoubleLiteral(ctx *DoubleLiteralContext) {}

// ExitDoubleLiteral is called when production doubleLiteral is exited.
func (s *BaseCypherListener) ExitDoubleLiteral(ctx *DoubleLiteralContext) {}

// EnterSchemaName is called when production schemaName is entered.
func (s *BaseCypherListener) EnterSchemaName(ctx *SchemaNameContext) {}

// ExitSchemaName is called when production schemaName is exited.
func (s *BaseCypherListener) ExitSchemaName(ctx *SchemaNameContext) {}

// EnterReservedWord is called when production reservedWord is entered.
func (s *BaseCypherListener) EnterReservedWord(ctx *ReservedWordContext) {}

// ExitReservedWord is called when production reservedWord is exited.
func (s *BaseCypherListener) ExitReservedWord(ctx *ReservedWordContext) {}

// EnterSymbolicName is called when production symbolicName is entered.
func (s *BaseCypherListener) EnterSymbolicName(ctx *SymbolicNameContext) {}

// ExitSymbolicName is called when production symbolicName is exited.
func (s *BaseCypherListener) ExitSymbolicName(ctx *SymbolicNameContext) {}

// EnterLeftArrowHead is called when production leftArrowHead is entered.
func (s *BaseCypherListener) EnterLeftArrowHead(ctx *LeftArrowHeadContext) {}

// ExitLeftArrowHead is called when production leftArrowHead is exited.
func (s *BaseCypherListener) ExitLeftArrowHead(ctx *LeftArrowHeadContext) {}

// EnterRightArrowHead is called when production rightArrowHead is entered.
func (s *BaseCypherListener) EnterRightArrowHead(ctx *RightArrowHeadContext) {}

// ExitRightArrowHead is called when production rightArrowHead is exited.
func (s *BaseCypherListener) ExitRightArrowHead(ctx *RightArrowHeadContext) {}

// EnterDash is called when production dash is entered.
func (s *BaseCypherListener) EnterDash(ctx *DashContext) {}

// ExitDash is called when production dash is exited.
func (s *BaseCypherListener) ExitDash(ctx *DashContext) {}
