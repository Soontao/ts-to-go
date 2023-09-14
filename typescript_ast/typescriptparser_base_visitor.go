// Code generated from ./lib/TypeScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package typescript_ast // TypeScriptParser
import "github.com/antlr4-go/antlr/v4"

type BaseTypeScriptParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTypeScriptParserVisitor) VisitInitializer(ctx *InitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBindingPattern(ctx *BindingPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeParameters(ctx *TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeParameterList(ctx *TypeParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitConstraint(ctx *ConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeArguments(ctx *TypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeArgumentList(ctx *TypeArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeArgument(ctx *TypeArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIntersection(ctx *IntersectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitUnion(ctx *UnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitRedefinitionOfType(ctx *RedefinitionOfTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPredefinedPrimType(ctx *PredefinedPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrayPrimType(ctx *ArrayPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitParenthesizedPrimType(ctx *ParenthesizedPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitThisPrimType(ctx *ThisPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTuplePrimType(ctx *TuplePrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitObjectPrimType(ctx *ObjectPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitReferencePrimType(ctx *ReferencePrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitQueryPrimType(ctx *QueryPrimTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPredefinedType(ctx *PredefinedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeReference(ctx *TypeReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitNestedTypeGeneric(ctx *NestedTypeGenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeGeneric(ctx *TypeGenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeIncludeGeneric(ctx *TypeIncludeGenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitObjectType(ctx *ObjectTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeBody(ctx *TypeBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeMemberList(ctx *TypeMemberListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeMember(ctx *TypeMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTupleType(ctx *TupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTupleElementTypes(ctx *TupleElementTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitConstructorType(ctx *ConstructorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeQuery(ctx *TypeQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeQueryExpression(ctx *TypeQueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertySignatur(ctx *PropertySignaturContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitCallSignature(ctx *CallSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitRequiredParameterList(ctx *RequiredParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitOptionalParameter(ctx *OptionalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitRestParameter(ctx *RestParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitRequiredParameter(ctx *RequiredParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAccessibilityModifier(ctx *AccessibilityModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIdentifierOrPattern(ctx *IdentifierOrPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitConstructSignature(ctx *ConstructSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIndexSignature(ctx *IndexSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMethodSignature(ctx *MethodSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeAliasDeclaration(ctx *TypeAliasDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitInterfaceExtendsClause(ctx *InterfaceExtendsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassOrInterfaceTypeList(ctx *ClassOrInterfaceTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEnumBody(ctx *EnumBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEnumMemberList(ctx *EnumMemberListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEnumMember(ctx *EnumMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitNamespaceDeclaration(ctx *NamespaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitNamespaceName(ctx *NamespaceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitImportAliasDeclaration(ctx *ImportAliasDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDecoratorList(ctx *DecoratorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDecorator(ctx *DecoratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDecoratorMemberExpression(ctx *DecoratorMemberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDecoratorCallExpression(ctx *DecoratorCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitSourceElement(ctx *SourceElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAbstractDeclaration(ctx *AbstractDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFromBlock(ctx *FromBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMultipleImportStatement(ctx *MultipleImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitVariableStatement(ctx *VariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDoStatement(ctx *DoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitForVarStatement(ctx *ForVarStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitForInStatement(ctx *ForInStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitForVarInStatement(ctx *ForVarInStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitVarModifier(ctx *VarModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitYieldStatement(ctx *YieldStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitWithStatement(ctx *WithStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitCaseClauses(ctx *CaseClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitCaseClause(ctx *CaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDefaultClause(ctx *DefaultClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitLabelledStatement(ctx *LabelledStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitThrowStatement(ctx *ThrowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTryStatement(ctx *TryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitCatchProduction(ctx *CatchProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFinallyProduction(ctx *FinallyProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDebuggerStatement(ctx *DebuggerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassHeritage(ctx *ClassHeritageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassTail(ctx *ClassTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassExtendsClause(ctx *ClassExtendsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitImplementsClause(ctx *ImplementsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitClassElement(ctx *ClassElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertyDeclarationExpression(ctx *PropertyDeclarationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMethodDeclarationExpression(ctx *MethodDeclarationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGetterSetterDeclarationExpression(ctx *GetterSetterDeclarationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAbstractMemberDeclaration(ctx *AbstractMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertyMemberBase(ctx *PropertyMemberBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIndexMemberDeclaration(ctx *IndexMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGeneratorMethod(ctx *GeneratorMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGeneratorFunctionDeclaration(ctx *GeneratorFunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGeneratorBlock(ctx *GeneratorBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGeneratorDefinition(ctx *GeneratorDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIteratorBlock(ctx *IteratorBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIteratorDefinition(ctx *IteratorDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFormalParameterArg(ctx *FormalParameterArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitLastFormalParameterArg(ctx *LastFormalParameterArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitSourceElements(ctx *SourceElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitElementList(ctx *ElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrayElement(ctx *ArrayElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertyGetter(ctx *PropertyGetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertySetter(ctx *PropertySetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMethodProperty(ctx *MethodPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertyShorthand(ctx *PropertyShorthandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitRestParameterInObject(ctx *RestParameterInObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGetAccessor(ctx *GetAccessorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitSetAccessor(ctx *SetAccessorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPropertyName(ctx *PropertyNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitExpressionSequence(ctx *ExpressionSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFunctionExpressionDeclaration(ctx *FunctionExpressionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTemplateStringExpression(ctx *TemplateStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTernaryExpression(ctx *TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGeneratorsExpression(ctx *GeneratorsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitInExpression(ctx *InExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGenericTypes(ctx *GenericTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArgumentsExpression(ctx *ArgumentsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitThisExpression(ctx *ThisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTypeofExpression(ctx *TypeofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitInstanceofExpression(ctx *InstanceofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitDeleteExpression(ctx *DeleteExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGeneratorsFunctionExpression(ctx *GeneratorsFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrowFunctionExpression(ctx *ArrowFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIteratorsExpression(ctx *IteratorsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBitXOrExpression(ctx *BitXOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitCastAsExpression(ctx *CastAsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitSuperExpression(ctx *SuperExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBitShiftExpression(ctx *BitShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitYieldExpression(ctx *YieldExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBitNotExpression(ctx *BitNotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitNewExpression(ctx *NewExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMemberDotExpression(ctx *MemberDotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitMemberIndexExpression(ctx *MemberIndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBitAndExpression(ctx *BitAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitBitOrExpression(ctx *BitOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitVoidExpression(ctx *VoidExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAsExpression(ctx *AsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrowFunctionDeclaration(ctx *ArrowFunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitArrowFunctionBody(ctx *ArrowFunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTemplateStringLiteral(ctx *TemplateStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitTemplateStringAtom(ctx *TemplateStringAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIdentifierName(ctx *IdentifierNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitIdentifierOrKeyWord(ctx *IdentifierOrKeyWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitReservedWord(ctx *ReservedWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitGetter(ctx *GetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitSetter(ctx *SetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypeScriptParserVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}
